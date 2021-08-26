package service

import (
	"strings"
)

func (bot *Bot) Handler(u *Update) error {
	msg := u.Message.Text
	chat := u.Message.Chat
	command := SanitizeCommand(msg)

	bot.log.Warn("Received something")
	if command == "/showEvents" {
		bot.ShowEvents(chat, strings.TrimPrefix(msg, command))
	}
	if command == "/addEvent" {
		bot.AddEvent(chat, strings.TrimPrefix(msg, command))
	}

	return nil
}
