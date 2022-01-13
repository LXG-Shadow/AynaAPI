package cmd

import (
	"AynaAPI/discord/app"
	"AynaAPI/discord/musicbot"
	"github.com/bwmarrin/discordgo"
)

func SongQuitCommand(ctx app.AppContext) {
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
	player.Stop()
	musicbot.DestroyMusicPlayer(ctx)
	ctx.InfoEmbedMessage(&discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title: "MusicBot Quit",
		},
	})
}
