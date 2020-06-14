package main

import (
	"fmt"
	"./config"
)

var BotID string
var goBot *dsciordgo.Session

func main() {
	err := config.ReadConfig()
	if err != nil{
		fmt Println(err.Error())
		return err
	}
	<-make(chan struct{})
	return
}
