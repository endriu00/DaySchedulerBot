package service

import (
	"errors"
)

var ErrIsBot = errors.New("User is a bot")

func (bot *Bot) AddUser(chatId int, user User) error {
	if user.IsBot {
		//Add a blacklist table in DB
		bot.log.Warn("Cannot add bots!")
		return ErrIsBot
	}
	_, err := bot.db.Exec(`INSERT INTO telegram_user(chatid, username) VALUES (?, ?)`, chatId, user.Username)
	if err != nil {
		bot.log.WithError(err).Error("Could not add user " + user.Username)
		return err
	}
	return nil
}
