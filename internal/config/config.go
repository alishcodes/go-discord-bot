package config

import "os"

type Config struct {
	APIEndpoint  string
	AIModel      string
	DiscordToken string
	BotTrigger   string
}

func Load() (*Config, error) {
	return &Config{
		APIEndpoint:  os.Getenv("API_ENDPOINT"),
		AIModel:      os.Getenv("AI_MODEL"),
		DiscordToken: os.Getenv("DISCORD_BOT_TOKEN"),
		BotTrigger:   os.Getenv("BOT_TRIGGER"),
	}, nil
}
