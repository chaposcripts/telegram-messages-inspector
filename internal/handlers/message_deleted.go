package handlers

import (
	"context"
	"fmt"
	"removed-messages/internal/config"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandleDeletedMessage(cfg *config.Config, ctx context.Context, b *bot.Bot, update *models.BusinessMessagesDeleted) {
	for _, deletedMessageID := range update.MessageIDs {
		deletedMessageText, exists := Messages[deletedMessageID]
		if !exists {
			deletedMessageText = &models.Message{Text: "<i>Unable to get the text of the deleted message :(</i>"}
		}

		username := update.Chat.Username
		if len(username) > 0 {
			username = fmt.Sprintf(" (@%s)", username)
		}

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    cfg.ChatID,
			ParseMode: "HTML",
			Text: fmt.Sprintf(`%s%s [<code>%d</code>] <b>delete message:</b>
<blockquote>%s</blockquote>`, update.Chat.FirstName, username, update.Chat.ID, deletedMessageText.Text),
		})
		b.SendMediaGroup(ctx, &bot.SendMediaGroupParams{})
	}
}
