package cmd

import (
	"AynaAPI/discord/app"
	"AynaAPI/discord/musicbot"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func SongQueueCommand(ctx app.AppContext) {
	player, err := musicbot.CreateMusicPlayer(ctx)
	if err != nil {
		ctx.ErrorEmbedMessage(&discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: err.Error(),
			},
		})
		return
	}
	var sb strings.Builder
	for index, rs := range player.UserPlaylist.Musics {
		if index >= 16 {
			break
		}
		sb.WriteString(fmt.Sprintf("#%d  %s - %s - %s\n", index, rs.Music.Title, rs.Music.Title, rs.User.Username))
	}
	ctx.InfoEmbedMessage(&discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "Current Queue",
			Description: sb.String(),
		},
	})
}
