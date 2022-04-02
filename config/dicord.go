package config

type Discord struct {
	BotToken string

	DefaultMusicProvider string
}

var DiscordConfig *Discord

func init() {
	DiscordConfig = &Discord{}
	if cfgFile == nil {
		return
	}
	mapTo(cfgFile, "Discord", DiscordConfig)
}
