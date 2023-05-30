package main

import (
	"fmt"
	"gopt/chatgpt"
	"gopt/discord"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var chatgptClient chatgpt.ChatgptClient
var discordBot *discordgo.Session
var prefix string

func main() {
	godotenv.Load()

	openaiToken, discordToken := getEnv("OPENAI_TOKEN"), getEnv("DISCORD_TOKEN")
	discordBot, chatgptClient = discord.NewBot(discordToken), chatgpt.NewClient(openaiToken)

	prefix = getEnv("PREFIX")
	discordBot.AddHandler(messageCreate)

	// whait...
	fmt.Scanln()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, prefix) {
		response := chatgptClient.GenerateResponse(m.Content[len(prefix):])
		s.ChannelMessageSend(m.ChannelID, response)
		return
	}
}

func getEnv(key string) string {
	key, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("Missing %s environment variable", key)
	}

	return key
}
