package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"

	"github.com/theovidal/carrbridge/commands"
	"github.com/theovidal/carrbridge/handlers"
)

func main() {
	bot := onyxcord.RegisterBot("carrbridge")

	bot.Client.AddHandler(func(_ *discordgo.Session, message *discordgo.MessageCreate) {
		handlers.MessageTransfer(&bot, message)
	})
	bot.RegisterCommand("plug", commands.Plug())

	bot.Client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	bot.Run(true)
}
