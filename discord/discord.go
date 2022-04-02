package discord

import (
	"AynaAPI/config"
	"AynaAPI/discord/app"
	"AynaAPI/discord/cmd"
	"github.com/bwmarrin/discordgo"
)

func CreateBot(token string) (*discordgo.Session, error) {
	if token == "" {
		token = config.DiscordConfig.BotToken
	}
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	registerCommand(session)
	return session, nil
}

func registerCommand(session *discordgo.Session) {
	session.AddHandler(app.DiscordCommandHandler(".info", cmd.SongInfoCommand))
	session.AddHandler(app.DiscordCommandHandler(".search", cmd.SongSearchCommand))
	session.AddHandler(app.DiscordCommandHandler(".play", cmd.SongPlayCommand))
	session.AddHandler(app.DiscordCommandHandler(".playnext", cmd.SongPlayNextCommand))
	session.AddHandler(app.DiscordCommandHandler(".quit", cmd.SongQuitCommand))
	session.AddHandler(app.DiscordCommandHandler(".pause", cmd.SongPauseCommand))
	session.AddHandler(app.DiscordCommandHandler(".unpause", cmd.SongUnpauseCommand))
}
