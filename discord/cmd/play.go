package cmd

import (
	"AynaAPI/discord/app"
	"AynaAPI/discord/musicbot"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func SongPlayCommand(ctx app.AppContext) {
	if len(ctx.Args) == 0 {
		ctx.MessageSend(fmt.Sprintf("Usage: %s <keywords>", ctx.Command))
		return
	}
	keyword := strings.Join(ctx.Args, " ")
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
	item, err := player.AddMusicByKeyword(keyword, ctx.GetAuthorAppUser())
	if err != nil {
		ctx.WarnEmbedMessage(&discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Description: fmt.Sprintf("Add fail: %s", err.Error()),
			},
		})
		return
	}
	ctx.InfoEmbedMessage(&discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "Music Added",
			Description: fmt.Sprintf("%s %s", item.Music.Title, item.Music.Artist),
		},
	})
	//ctx.Logger().Info(sendComplex)
	//ctx.Logger().Info(err)
}
