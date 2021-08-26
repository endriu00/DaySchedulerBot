package service

import (
	"net/http"
)

func (bot *Bot) HandleTelegramWebhook(r *http.Request) {

	update, err := ParseTelegramRequest(r)
	if err != nil {
		return
	}
	bot.Handler(update)
}
