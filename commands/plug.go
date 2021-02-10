package commands

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"

	"github.com/theovidal/carrbridge/lib"
)

func Plug() *onyxcord.Command {
	return &onyxcord.Command{
		Description:    "Ajouter un salon dans un routeur",
		Usage:          "plug <token>",
		Show:           true,
		ListenInPublic: true,
		Execute: func(arguments []string, bot *onyxcord.Bot, message *discordgo.MessageCreate) (err error) {
			bot.Client.ChannelMessageDelete(message.ChannelID, message.ID)

			var router lib.Router
			router, err = lib.GetRouterFromToken(bot, arguments[0])
			if err != nil {
				return errors.New("Le routeur ciblé est inconnu. Essayez d'utiliser une autre clé.")
			}

			err = lib.CreatePlug(bot, &router, message.ChannelID)
			if err != nil {
				return
			}

			_, err = bot.Client.ChannelMessageSendEmbed(message.ChannelID, onyxcord.MakeEmbed(
				bot.Config,
				&discordgo.MessageEmbed{
					Title:       fmt.Sprintf("🔌 Ce salon est désormais connecté au routeur `%s`!", router.Name),
					Description: "Tous les messages qui y seront envoyés seront transférés sur les autres salons connectés, et inversement.",
				}),
			)
			return
		},
	}
}
