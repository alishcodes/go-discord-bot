package bot

import (
	"github.com/alishcodes/go-discord-bot/internal/aliceapi"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session
	AI      *aliceapi.Client
	Trigger string
}

func New(s *discordgo.Session, client *aliceapi.Client, t string) *Bot {
	return &Bot{
		Session: s,
		AI:      client,
		Trigger: t,
	}
}
