package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Token  string
	ChatID uint
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	token := os.Getenv("BOT_TOKEN")
	chatID, err := strconv.ParseUint(os.Getenv("CHAT_ID"), 10, 64)
	if err != nil {
		return nil, err
	}

	return &Config{
		Token:  token,
		ChatID: uint(chatID),
	}, nil
}
