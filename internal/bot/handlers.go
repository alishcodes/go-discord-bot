package bot

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) RegisterHandlers() {
	b.Session.AddHandler(b.onMessage)
}

func (b *Bot) onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.Author.ID == s.State.SessionID {
		return
	}

	emojis, err := s.GuildEmojis(m.GuildID)
	if err != nil {
		log.Fatal("error loading emojis:", err)
	}

	content := strings.ToLower(m.Content)

	if strings.Contains(content, b.Trigger) {
		// 30% chance for bot to react to a message
		if rand.Intn(100) < 30 {
			rn := rand.Intn(len(emojis)) // select rand num between 0 - num of emojis
			emoji := emojis[rn]

			var emojiID string
			if emoji.Animated {
				emojiID = fmt.Sprintf(`a:%s:%s`, emoji.Name, emoji.ID)
			} else {
				emojiID = fmt.Sprintf(`%s:%s`, emoji.Name, emoji.ID)
			}

			s.MessageReactionAdd(m.ChannelID, m.ID, emojiID) // Add random emoji reaction to users message
		}

		// Format prompt for Alice api
		prompt := fmt.Sprintf(
			"Current speaker: %s. Message: %s",
			m.Author.DisplayName(),
			m.Content,
		)

		log.Println("Sending request")

		// make channel to stop typing once response is received.
		typingChan := make(chan struct{})

		// start typing while waiting for response
		go b.setTyping(m.ChannelID, typingChan)

		go b.generateResponse(m.Reference(), m.ChannelID, prompt, typingChan)
	}
}

func (b *Bot) generateResponse(m *discordgo.MessageReference, channelID string, prompt string, typingChan chan struct{}) {
	response := b.AI.SendRequest(prompt)

	log.Println("Received response")

	close(typingChan)

	//do not send blank response.
	if response != "" {
		b.Session.ChannelMessageSendReply(channelID, response, m)
	}
}
