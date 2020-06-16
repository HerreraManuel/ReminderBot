package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/HerreraManuel/ReminderBot/config"
	"github.com/bwmarrin/discordgo"
)

func main() {
	config.ReadConfig()

	// Creating discord session with parsed token
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Error creating discord session,", err)
		return
	}

	// Register the message create func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	// Let bot run and indicate to user how to exit
	fmt.Println("Bot is running...Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close the discord session
	dg.Close()
}

// Whenever a new message is receieved, this func will be called.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Bot checking if prefix is called
	if strings.HasPrefix(m.Content, config.BotPrefix) {

		// Ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		// TESTING WITH PING PONG
		if m.Content == "!ping" {
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		}

		if m.Content == "!pong" {
			s.ChannelMessageSend(m.ChannelID, "Ping!")
		}

		if m.Content == "!remind" {
			s.ChannelMessageSend(m.ChannelID, "placeholder")
		}
	}
}
