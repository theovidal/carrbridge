package handlers

import (
	"encoding/hex"
	"log"
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"

	"github.com/theovidal/carrbridge/lib"
)

func _() string {
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

// MessageTransfer handles new messages inside channels linked in a router, and transfer them upon the router
func MessageTransfer(bot *onyxcord.Bot, msg *discordgo.MessageCreate) {
	if msg.Author.Bot {
		return
	}

	routers, err := lib.GetRoutersFromChannel(bot, msg.ChannelID)
	if err != nil {
		log.Println(err)
		return
	}

	for _, router := range routers {
		for _, plug := range router.Plugs {
			if plug.Channel == msg.ChannelID {
				continue
			}

			avatar := msg.Author.AvatarURL("128")
			var name string
			member, _ := bot.Client.GuildMember(msg.GuildID, msg.Author.ID)

			if member.Nick == "" {
				name = msg.Author.Username
			} else {
				name = member.Nick
			}

			content := msg.Content
			for _, file := range msg.Attachments {
				content += file.ProxyURL
			}

			if content != "" {
				_, err := bot.Client.WebhookExecute(plug.Webhook, plug.Token, false, &discordgo.WebhookParams{
					Username:  name,
					AvatarURL: avatar,
					Content:   content,
				})
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
