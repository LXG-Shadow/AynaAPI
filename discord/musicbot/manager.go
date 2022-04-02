package musicbot

import (
	"AynaAPI/discord/app"
)

var MusicPlayerList map[string]*MusicPlayer

func init() {
	MusicPlayerList = map[string]*MusicPlayer{}
}

func CreateMusicPlayer(ctx app.AppContext) (*MusicPlayer, error) {
	p := GetMusicPlayerByGuildID(ctx.GetGuild().ID)
	if p != nil {
		return p, nil
	}
	vc, err := ctx.GetVoiceChannel()
	if err != nil {
		return nil, err
	}
	vcc, err := ctx.Session.ChannelVoiceJoin(ctx.GetGuild().ID, vc.ID, false, false)
	if err != nil {
		return nil, err
	}
	p = new_MusicPlayer(ctx.GetGuild().ID, vc.ID, vcc)
	MusicPlayerList[ctx.GetGuild().ID] = p
	return p, nil
}

func DestroyMusicPlayer(ctx app.AppContext) {
	guild := ctx.GetGuild()
	p := GetMusicPlayerByGuildID(guild.ID)
	if p == nil {
		return
	}
	p.Close()
	MusicPlayerList[guild.ID].VCConnection.Disconnect()
	delete(MusicPlayerList, guild.ID)
}

func GetMusicPlayerByGuildID(guildID string) *MusicPlayer {
	p, ok := MusicPlayerList[guildID]
	if !ok {
		return nil
	}
	return p
}
