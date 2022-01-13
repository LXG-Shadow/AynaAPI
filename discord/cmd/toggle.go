package cmd

import (
	"AynaAPI/discord/app"
	"AynaAPI/discord/musicbot"
	"github.com/bwmarrin/discordgo"
)

func SongPauseCommand(ctx app.AppContext) {
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
	err = player.Toggle(true)
	if err != nil {
		return
	}
	ctx.InfoEmbedMessage(&discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title: "Music Paused",
		},
	})
}

func SongUnpauseCommand(ctx app.AppContext) {
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
	err = player.Toggle(false)
	if err != nil {
		return
	}
	ctx.InfoEmbedMessage(&discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title: "Music Unpaused",
		},
	})
}
