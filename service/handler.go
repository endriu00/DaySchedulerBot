package service

import (
	"strings"
)

func Handler(u *Update) error {
	msg := u.Message.Text
	chat := u.Message.Chat
	command := SanitizeCommand(msg)
	if command == "/showEvents" {
		ShowEvents(chat, strings.TrimPrefix(msg, command))
	}
	if command == "/addEvent" {
		AddEvent(chat, strings.TrimPrefix(msg, command))
	}

	return nil
}
