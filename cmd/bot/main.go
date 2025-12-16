package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alishcodes/go-discord-bot/internal/aliceapi"
	"github.com/alishcodes/go-discord-bot/internal/bot"
	"github.com/alishcodes/go-discord-bot/internal/config"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file **FOR LOCAL DEV**
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using system env variables.")
	}

	// load env variables to config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("error loading .env config:", err)
	}

	discord, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Fatal("error creating new discord session: ", err)
	}

	discord.Identify.Intents =
		discordgo.IntentGuildMessages |
			discordgo.IntentMessageContent

	//start new aliceapi client
	alice := aliceapi.New(cfg.APIEndpoint, cfg.AIModel)

	//create alicebot with discord session and aliceapi client
	aliceBot := bot.New(discord, alice, cfg.BotTrigger)
	aliceBot.RegisterHandlers()

	err = discord.Open()
	if err != nil {
		log.Fatal("error opening connection:", err)
	}

	fmt.Println("The bot is online! Press Ctrl + C to shut down bot.")

	sc := make(chan os.Signal, 1)
	//wait for ctrl+c to be hit then signal sent to channel and program finishes
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sc

	fmt.Println("Shutting down bot...")

	discord.Close()
}
