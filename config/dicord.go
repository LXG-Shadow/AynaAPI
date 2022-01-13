package config

type Discord struct {
	BotToken string

	DefaultMusicProvider string
}

var DiscordConfig *Discord

func init() {
	DiscordConfig = &Discord{}
	mapTo(cfgFile, "Discord", DiscordConfig)
}
