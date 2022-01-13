package cmd

import (
	"AynaAPI/discord/app"
	"AynaAPI/discord/musicbot"
	"github.com/bwmarrin/discordgo"
)

func SongPlayNextCommand(ctx app.AppContext) {

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
	player.PlayNext()
	//if err != nil {
	//	ctx.WarnEmbedMessage(&discordgo.MessageSend{
	//		Embed: &discordgo.MessageEmbed{
	//			Description: err.Error(),
	//		},
	//	})
	//	return
	//}
	//ctx.InfoEmbedMessage(&discordgo.MessageSend{
	//	Embed: &discordgo.MessageEmbed{
	//		Title:       "Music Playing",
	//		Description: fmt.Sprintf("%s %s", item.Music.Title, item.Music.Artist),
	//	},
	//})
}
