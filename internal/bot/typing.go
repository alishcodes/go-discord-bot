package bot

import (
	"log"
	"time"
)

func (b *Bot) setTyping(channelID string, stopChan chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)

	defer ticker.Stop()

	err := b.Session.ChannelTyping(channelID)
	if err != nil {
		log.Println("error setting typing status:", err)
		return
	}

	for {
		select {
		case <-ticker.C: // Set typing status every 5 seconds while waiting for response
			err = b.Session.ChannelTyping(channelID)
			if err != nil {
				log.Println("error setting typing status:", err)
				return
			}
		case <-stopChan:
			return
		}
	}
}
