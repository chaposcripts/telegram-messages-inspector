package handlers

import (
	"context"
	"fmt"
	"removed-messages/internal/config"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandleEditedMessage(cfg *config.Config, ctx context.Context, b *bot.Bot, update *models.Message) {
	originalMessageText, exists := Messages[update.ID]
	if !exists {
		originalMessageText = &models.Message{Text: "<i>Unable to get the text of the edited message :(</i>"}
	}

	username := update.Chat.Username
	if len(username) > 0 {
		username = " @" + username
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    cfg.ChatID,
		ParseMode: "HTML",
		Text: fmt.Sprintf(`%s%s [<code>%d</code>] <b>edited message:</b>
<i>Before:</i>
<blockquote>%s</blockquote>
<i>After:</i>.
<blockquote>%s</blockquote>`,
			update.Chat.FirstName, username, update.Chat.ID, originalMessageText.Text, update.Text),
	})
	if exists {
		Messages[update.ID].Text = fmt.Sprintf("<s>%s</s> %s", originalMessageText.Text, update.Text)
	}
}
