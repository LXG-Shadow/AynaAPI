package musicbot

import (
	"AynaAPI/api/music"
	"AynaAPI/config"
	"AynaAPI/discord/app"
	"AynaAPI/discord/e"
	"AynaAPI/discord/pkg/dgvoice"
	"github.com/bwmarrin/discordgo"
	"sync"
)

type MusicItem struct {
	Music music.Music
	User  app.AppUser
}

type Playlist struct {
	Index  int
	Musics []MusicItem
}

func (p *Playlist) Clear() {
	p.Musics = nil
}

func (p *Playlist) IsEmpty() bool {
	return len(p.Musics) == 0
}

func (p *Playlist) PopFirst() (item *MusicItem) {
	if p.IsEmpty() {
		return nil
	}
	item, p.Musics = &p.Musics[0], p.Musics[1:]
	return
}

func (p *Playlist) Next() (item *MusicItem) {
	if p.IsEmpty() {
		return
	}
	if p.Index >= len(p.Musics) {
		p.Index = 0
	}
	item = &p.Musics[p.Index]
	p.Index++
	return
}

func (p *Playlist) Add(item MusicItem) {
	p.Musics = append(p.Musics, item)
}

type MusicPlayer struct {
	GuildID      string
	VCChannelID  string
	VCConnection *discordgo.VoiceConnection

	UserPlaylist   Playlist
	SystemPlaylist Playlist
	CurrentMusic   *MusicItem
	Paused         bool

	StopChan  chan bool
	PauseChan chan bool

	mu sync.Mutex
}

func new_MusicPlayer(GuildID, VCChannelID string, VCConnection *discordgo.VoiceConnection) *MusicPlayer {
	return &MusicPlayer{
		GuildID:        GuildID,
		VCChannelID:    VCChannelID,
		VCConnection:   VCConnection,
		UserPlaylist:   Playlist{},
		SystemPlaylist: Playlist{},

		StopChan:  make(chan bool),
		PauseChan: make(chan bool),
	}
}

func (m *MusicPlayer) IsPlaying() bool {
	return m.CurrentMusic != nil
}

func (m *MusicPlayer) Play(item *MusicItem) error {
	audio, errcode := MusicGetAudioFromMusic(item.Music)
	if errcode != 0 {
		return e.NewError(errcode)
	}
	m.mu.Lock()
	m.CurrentMusic = item
	m.Paused = false
	m.mu.Unlock()
	go func() {
		m.mu.Lock()
		dgvoice.PlayAudioFile(m.VCConnection, audio.Url, m.StopChan, m.PauseChan)
		m.CurrentMusic = nil
		m.StopChan = make(chan bool)
		m.Paused = true
		m.PauseChan = make(chan bool)
		m.mu.Unlock()
		defer m.playNext()
	}()
	return nil
}

func (m *MusicPlayer) PlayNext() {
	m.Stop()
	return
}

func (m *MusicPlayer) playNext() (*MusicItem, error) {
	var item *MusicItem
	if !m.UserPlaylist.IsEmpty() {
		item = m.UserPlaylist.PopFirst()
	} else {
		if !m.SystemPlaylist.IsEmpty() {
			item = m.SystemPlaylist.Next()
		}
	}
	if item == nil {
		return item, e.NewError(e.MUSICBOT_ERROR_QUEUE_EMPTY)
	}
	if m.IsPlaying() {
		m.Stop()
	}
	return item, m.Play(item)
}

func (m *MusicPlayer) Toggle(pause bool) error {
	if !m.IsPlaying() {
		return e.NewError(e.MUSICBOT_ERROR_ALREADY_PAUSED_UNPAUSED)
	}
	if m.Paused == pause {
		return e.NewError(e.MUSICBOT_ERROR_ALREADY_PAUSED_UNPAUSED)
	}
	m.Paused = pause
	m.PauseChan <- true
	return nil
}

func (m *MusicPlayer) Stop() {
	if m.IsPlaying() {
		m.StopChan <- true
	}
}

func (m *MusicPlayer) Close() {
	m.mu.Lock()
	m.Stop()
	m.mu.Unlock()
}

func (m *MusicPlayer) AddMusicItem(item MusicItem) (MusicItem, error) {
	m.UserPlaylist.Add(item)
	if !m.IsPlaying() {
		defer m.playNext()
	}
	return item, nil
}

func (m *MusicPlayer) AddMusicByMetadata(metadata string, user app.AppUser) (MusicItem, error) {
	musicc, errcode := MusicGet(metadata)
	if errcode != 0 {
		return MusicItem{}, e.NewError(errcode)
	}
	return m.AddMusicItem(MusicItem{Music: musicc, User: user})
}

func (m *MusicPlayer) AddMusicByKeyword(keyword string, user app.AppUser) (MusicItem, error) {
	result, errcode := MusicSearch(config.DiscordConfig.DefaultMusicProvider, keyword)
	if errcode != 0 {
		return MusicItem{}, e.NewError(errcode)
	}
	if len(result.Result) == 0 {
		return MusicItem{}, e.NewError(e.MUSIC_ERROR_SEARCH_FAIL)
	}
	musicc, errcode := MusicGetFromMusicMeta(result.Result[0])
	if errcode != 0 {
		return MusicItem{}, e.NewError(errcode)
	}
	return m.AddMusicItem(MusicItem{Music: musicc, User: user})
}
