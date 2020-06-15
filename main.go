package main

import (
	"fmt"

	"github.com/HerreraManuel/ReminderBot/config"
)

func main() {

	config.ReadConfig()

	fmt.Println("Token:", configItems.Token)

	fmt.Println("ReminderBot is running...CTRL+C to exit")

}
