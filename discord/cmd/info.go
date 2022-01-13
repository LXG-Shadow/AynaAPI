package cmd

import (
	"AynaAPI/config"
	"AynaAPI/discord/app"
	"AynaAPI/discord/e"
	"AynaAPI/discord/musicbot"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func SongInfoCommand(ctx app.AppContext) {
	if len(ctx.Args) == 0 {
		ctx.MessageSend(fmt.Sprintf("Usage: %s <keywords>", ctx.Command))
		return
	}
	keyword := strings.Join(ctx.Args, " ")
	result, errcode := musicbot.MusicSearch(config.DiscordConfig.DefaultMusicProvider, keyword)
	if errcode != e.SUCCESS {
		ctx.MessageSend(fmt.Sprintf("Search fail: %s", e.NewError(errcode)))
	}
	if len(result.Result) == 0 {
		ctx.MessageSend("no search result")
		return
	}
	song := result.Result[0]
	ctx.InfoEmbedMessage(&discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       song.Title,
			Description: song.Artist,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Source: " + song.Provider.Name,
			},
			Image: &discordgo.MessageEmbedImage{
				URL:    result.Result[0].Cover,
				Width:  128,
				Height: 128,
			},
		},
	})
	//ctx.Logger().Info(sendComplex)
	//ctx.Logger().Info(err)
}
