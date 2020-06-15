package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// variables that will be stored
var (
	Token     string
	BotPrefix string

	configItems *configStruct
)

type configStruct struct {
	Token     string
	BotPrefix string
}

// ReadConfig parses the json file to read and stored
// the bot token and the prefix for the bot to respond
func ReadConfig() {
	fmt.Println("Reading json file...")
	jsonFile, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		fmt.Println("error:", err)
	}
	err = json.Unmarshal(jsonFile, &configItems)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Storing credentials...")
	// Assigning the parsed json
	Token = configItems.Token
	BotPrefix = configItems.BotPrefix

	return
}
