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

func SongSearchCommand(ctx app.AppContext) {
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
	var sb strings.Builder
	for index, rs := range result.Result {
		if index >= 16 {
			break
		}
		sb.WriteString(fmt.Sprintf("#%d  %s - %s\n", index, rs.Title, rs.Artist))
	}
	ctx.InfoEmbedMessage(&discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("\"%s\" Result", keyword),
			Description: sb.String(),
		},
	})
}
