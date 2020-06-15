package main

import (
	"fmt"

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

}
