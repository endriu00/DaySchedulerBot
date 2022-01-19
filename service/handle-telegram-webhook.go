package service

import (
	"net/http"
)

func (bot *Bot) HandleTelegramWebhook(w http.ResponseWriter, r *http.Request) {

	update, err := ParseTelegramRequest(r)
	if err != nil {
		bot.log.WithError(err).Error("Cannot parse telegram request.")
		return
	}
	err = bot.Handler(update)
	if err != nil {
		bot.log.WithError(err).Error("Handler returned an error.")
		return
	}
}
