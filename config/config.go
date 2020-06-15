package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token string
	BotPrefix
	configItems *configStruct
)

type configStruct struct {
	Token     string
	BotPrefix string
}

func ReadConfig() {
	fmt.Println("Opening File...")
	jsonFile, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		fmt.Println("error:", err)
	}
	err = json.Unmarshal(jsonFile, &configItems)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Token:", configItems.Token)
	fmt.Println("BotPrefix:", configItems.BotPrefix)
}
