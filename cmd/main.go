package main

import (
	"log"
	"removed-messages/internal/bot"
	"removed-messages/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = bot.Init(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
}
