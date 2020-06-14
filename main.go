package main

import (
	"fmt"
	"./config"
)

var BotID string
var goBot *discordGo.Session

func main() {
	err := config.ReadConfig()
	if err != nil{
		fmt Println(err.Error())
		return err
	}
	Start()
	<-make(chan struct{})
	return
}

func Start(){
	goBot, err := discordgo.New("Bot " + config.Token)
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@Me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running")
}
