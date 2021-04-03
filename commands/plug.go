package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"

	"github.com/theovidal/carrbridge/lib"
)

func Plug() *onyxcord.Command {
	return &onyxcord.Command{
		ListenInPublic: true,
		Execute: func(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) (err error) {
			var router lib.Router
			router, err = lib.GetRouterFromToken(bot, interaction.Data.Options[0].StringValue())
			if err != nil {
				return bot.UserError(interaction, "Le routeur ciblé est inconnu. Essayez d'utiliser une autre clé.")
			}

			err = lib.CreatePlug(bot, &router, interaction.ChannelID)
			if err != nil {
				return
			}

			_ = bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Embeds: []*discordgo.MessageEmbed{
						bot.MakeEmbed(&discordgo.MessageEmbed{
							Title:       fmt.Sprintf("🔌 Ce salon est désormais connecté au routeur `%s`!", router.Name),
							Description: "Tous les messages qui y seront envoyés seront transférés sur les autres salons connectés, et inversement.",
						}),
					},
				},
			})
			return
		},
	}
}
