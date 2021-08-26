package service

import (
	"time"
)

func (bot *Bot) PutEvent(eventTime time.Time, eventDescription, username string) error {
	_, err := bot.db.Exec(`INSERT INTO event(username, time, description) VALUES (?, ?, ?)`, username, eventTime, eventDescription)
	if err != nil {
		bot.log.WithError(err).Error("Could not add an event for user " + username)
		return err
	}
	return nil
}
