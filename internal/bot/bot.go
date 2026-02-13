package bot

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	// "strings"

	"removed-messages/internal/config"
	"removed-messages/internal/handlers"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Init(cfg *config.Config) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(func(ctx context.Context, b *bot.Bot, update *models.Update) {
			if update.BusinessMessage != nil {
				if uint(update.BusinessMessage.Chat.ID) != cfg.ChatID {
					handlers.HandleMessage(ctx, b, update.BusinessMessage)
				}
			} else if update.DeletedBusinessMessages != nil {
				handlers.HandleDeletedMessage(cfg, ctx, b, update.DeletedBusinessMessages)
			} else if update.EditedBusinessMessage != nil {
				handlers.HandleEditedMessage(cfg, ctx, b, update.EditedBusinessMessage)
			} else if update.Message != nil {
				if update.Message.Text == "/start" || update.Message.Text == "/id" {
					b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID:    cfg.ChatID,
						ParseMode: "HTML",
						Text:      fmt.Sprintf("Your chat ID: <code>%d</code>", update.Message.Chat.ID),
					})
				}
			}
		}),
	}

	b, err := bot.New(cfg.Token, opts...)
	if err != nil {
		return err
	}
	log.Println("Bot started!")
	b.Start(ctx)
	return nil
}
