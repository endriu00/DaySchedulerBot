package service

import (
	"database/sql"
)

func (bot *Bot) GetUsername(chatid int) (string, error) {
	var username string
	rowUsername := bot.db.QueryRow(`SELECT username FROM telegram_user WHERE chatid=?`, chatid)
	err := rowUsername.Scan(&username)
	if err == sql.ErrNoRows {
		bot.log.WithError(err).Error("Could not find this username!")
		return "", err
	}
	if err != nil {
		bot.log.WithError(err).Error("An error occurred querying for username")
		return "", err
	}
	return username, nil
}
