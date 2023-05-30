package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func NewBot(token string) *discordgo.Session {
	bot, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatalln(err.Error())
	}

	bot.Identify.Intents = discordgo.IntentsGuildMessages
	bot.Open()

	log.Print("Discord bot created successfully")
	return bot
}
