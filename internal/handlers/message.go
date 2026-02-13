package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var Messages = map[int]*models.Message{}

func HandleMessage(ctx context.Context, b *bot.Bot, update *models.Message) {
	Messages[update.ID] = update
}
