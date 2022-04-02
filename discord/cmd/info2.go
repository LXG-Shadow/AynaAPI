package cmd

//
//import (
//	"AynaAPI/config"
//	"AynaAPI/discord/app"
//	"AynaAPI/discord/e"
//	"AynaAPI/discord/musicbot"
//	"AynaAPI/utils/vfile"
//	"bytes"
//	"fmt"
//	"github.com/bwmarrin/discordgo"
//	"github.com/go-resty/resty/v2"
//	"strings"
//)
//
//func SongInfoCommand(ctx app.AppContext) {
//	if len(ctx.Args) == 0 {
//		ctx.MessageSend(fmt.Sprintf("Usage: %s <keywords>", ctx.Command))
//		return
//	}
//	keyword := strings.Join(ctx.Args, " ")
//	result, errcode := musicbot.MusicSearch(config.DiscordConfig.DefaultMusicProvider, keyword)
//	if errcode != e.SUCCESS {
//		ctx.MessageSend(fmt.Sprintf("Search fail: %s", e.NewError(errcode)))
//	}
//	if len(result.Result) == 0 {
//		ctx.MessageSend("no search result")
//		return
//	}
//	ctx.Logger().Info(result.Result[0].Cover)
//	i, _ := resty.New().R().Get(result.Result[0].Cover)
//	sendComplex, err := ctx.Session.ChannelMessageSendComplex(ctx.Message.ChannelID, &discordgo.MessageSend{
//		File: &discordgo.File{
//			Name:        result.Result[0].Cover,
//			ContentType: "image/" + vfile.GetFileExt(result.Result[0].Cover),
//			Reader:      bytes.NewReader(i.Body()), //so me type that implements os.Reader holding png file data. Ex: io.*bytes.Buffer
//		},
//		Embed: &discordgo.MessageEmbed{
//			URL:         "",
//			Type:        discordgo.EmbedTypeImage,
//			Title:       "asdf",
//			Description: "asdfffff",
//			Timestamp:   "",
//			Color:       0,
//			Footer:      nil,
//			Image: &discordgo.MessageEmbedImage{
//				URL:      result.Result[0].Cover,
//				ProxyURL: "",
//				Width:    0,
//				Height:   0,
//			},
//			Thumbnail: nil,
//			Video:     nil,
//			Provider:  nil,
//			Author:    nil,
//			Fields:    nil,
//		},
//	})
//	ctx.Logger().Info(sendComplex)
//	ctx.Logger().Info(err)
//}
