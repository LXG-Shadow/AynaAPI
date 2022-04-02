package app

import (
	"AynaAPI/discord/e"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type AppContext struct {
	Session *discordgo.Session
	Message *discordgo.MessageCreate

	Command string
	Args    []string
}

func (ctx *AppContext) GetTextChannel() *discordgo.Channel {
	c, _ := ctx.Session.State.Channel(ctx.Message.ChannelID)
	return c
}

func (ctx *AppContext) GetVoiceChannel() (*discordgo.Channel, error) {
	for _, state := range ctx.GetGuild().VoiceStates {
		if state.UserID == ctx.GetSender().ID {
			channel, _ := ctx.Session.State.Channel(state.ChannelID)
			return channel, nil
		}
	}
	return nil, e.NewError(e.CONTEXT_ERROR_NOT_IN_VOICE_CHANNEL)
}

func (ctx *AppContext) GetGuild() *discordgo.Guild {
	g, _ := ctx.Session.State.Guild(ctx.Message.GuildID)
	return g
}

func (ctx *AppContext) GetSender() *discordgo.User {
	return ctx.Message.Author
}

func (ctx *AppContext) GetAuthorAppUser() AppUser {
	return AppUser{UserID: ctx.Message.Author.ID, Username: ctx.Message.Author.Username}
}

func (ctx *AppContext) MessageSend(content string) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, content)
}

func (ctx *AppContext) InfoEmbedMessage(message *discordgo.MessageSend) (*discordgo.Message, error) {
	message.Embed.Type = discordgo.EmbedTypeRich
	message.Embed.Color = 0x2ECC71
	return ctx.Session.ChannelMessageSendComplex(ctx.Message.ChannelID, message)
}

func (ctx *AppContext) WarnEmbedMessage(message *discordgo.MessageSend) (*discordgo.Message, error) {
	message.Embed.Type = discordgo.EmbedTypeRich
	message.Embed.Color = 0xF1C40F
	return ctx.Session.ChannelMessageSendComplex(ctx.Message.ChannelID, message)
}

func (ctx *AppContext) ErrorEmbedMessage(message *discordgo.MessageSend) (*discordgo.Message, error) {
	message.Embed.Type = discordgo.EmbedTypeRich
	message.Embed.Color = 0xCB4335
	return ctx.Session.ChannelMessageSendComplex(ctx.Message.ChannelID, message)
}

func (ctx *AppContext) Logger() *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"command": ctx.Command,
		"args":    ctx.Args,
		"guild":   fmt.Sprintf("%s(%s)", ctx.GetGuild().Name, ctx.GetGuild().ID),
		"channel": fmt.Sprintf("%s(%s)", ctx.GetTextChannel().Name, ctx.GetTextChannel().ID),
		"sender":  ctx.GetSender(),
	})
}

type AppUser struct {
	UserID   string
	Username string
}
