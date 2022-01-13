package app

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type discordHandler func(s AppContext)

func DiscordCommandHandler(command string, handler discordHandler) interface{} {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}
		if strings.HasPrefix(m.Content, command+" ") || m.Content == command {
			ss := strings.FieldsFunc(strings.TrimPrefix(m.Content, command), func(r rune) bool {
				return r == ' '
			})
			handler(AppContext{Session: s, Message: m, Command: command, Args: ss})
		}
	}
}
