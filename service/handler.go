package service

import (
	"database/sql"
	"strings"
)

func (bot *Bot) Handler(u *Update) error {
	var err error
	msg := u.Message.Text
	chat := u.Message.Chat
	command := SanitizeCommand(msg)

	if command == "/start" {
		_, err := bot.GetUsername(chat.Id)
		if err == sql.ErrNoRows {
			err = bot.AddUser(chat.Id, u.Message.From)
			if err == ErrIsBot {
				bot.log.WithError(err).Error("User was a bot")
				return err
			}
			bot.log.Info("User " + u.Message.From.Username + "is a new entry!")
		}
		if err != nil && err != sql.ErrNoRows {
			bot.log.WithError(err).Error("Failed getting username: Handler")
		}
	}
	if command == "/showEvents" {
		err = bot.ShowEvents(chat, strings.TrimPrefix(msg, command))
		if err != nil {
			bot.log.WithError(err).Error("Cannot show events")
			return err
		}
	}
	if command == "/addEvent" {
		err = bot.AddEvent(u, strings.TrimPrefix(msg, command))
		if err != nil {
			bot.log.WithError(err).Error("Cannot add event")
			return err
		}
	}

	return nil
}
