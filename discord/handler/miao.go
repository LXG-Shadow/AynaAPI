package handler

//
//import (
//	"AynaAPI/discord/app"
//	"fmt"
//)
//
//// This function will be called (due to AddHandler above) every time a new
//// message is created on any channel that the autenticated bot has access to.
//func _messageCreate(s *app.AppContext) {
//	c, err := s.GetTextChannel()
//	if err != nil {
//		// Could not find channel.
//		return
//	}
//	fmt.Println(c)
//	// Find the guild for that channel.
//	g, err := s.GetGuild()
//	if err != nil {
//		// Could not find guild.
//		return
//	}
//	fmt.Println(g)
//}
//
//var MessageCreate = app.DiscordCommandHandler("!test", _messageCreate)
//
//func _autoreply(s *app.AppContext) {
//	c, err := s.GetTextChannel()
//	if err != nil {
//		// Could not find channel.
//		return
//	}
//	s.Session.ChannelMessageSend(c.ID, "ybb")
//	// Find the guild for that channel.
//	g, err := s.GetGuild()
//	if err != nil {
//		// Could not find guild.
//		return
//	}
//	fmt.Println(g)
//	s.Session.ChannelVoiceJoin()
//}
//
//var Autoreply = app.DiscordCommandHandler("ybb", _autoreply)
