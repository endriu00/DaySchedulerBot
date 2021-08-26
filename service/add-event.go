package service

import (
	"bufio"
	"bytes"
	"database/sql"
	"strings"
	"time"
)

func (bot *Bot) AddEvent(update *Update, messageText string) error {
	//message should be in the form [from, to, what] for example [2021-08-29T17:00-2021-08-29T18:00, meeting with someone at somewhere]
	message := update.Message
	var eventTime time.Time
	var eventDescription string
	var err error
	scanner := bufio.NewScanner(strings.NewReader(messageText))
	splitCommas := func(data []byte, isEOF bool) (advance int, token []byte, err error) {
		commaIndex := bytes.IndexByte(data, ',')
		if commaIndex > 0 {
			// we need to return the next position
			buffer := data[:commaIndex]
			return commaIndex + 1, bytes.TrimSpace(buffer), nil
		}

		if isEOF {
			if len(data) > 0 {
				return len(data), bytes.TrimSpace(data), nil
			}
		}
		return 0, nil, nil
	}
	scanner.Split(splitCommas)

	username, err := bot.GetUsername(message.Chat.Id)
	if err == sql.ErrNoRows {
		bot.log.WithError(err).Error("Unable to find username for this chat id. Aborting putting event.")
		return err
	}
	if err != nil && err != sql.ErrNoRows {
		bot.log.WithError(err).Error("Error querying for username. Aborting putting event.")
		return err
	}
	index := 0
	for scanner.Scan() {
		if index == 0 {
			eventTime, err = time.Parse("2006-02-01 15:04", scanner.Text())
			if err != nil {
				bot.log.WithError(err).Error("Could not parse string given as input by user " + username)
				return err
			}
		}
		if index == 1 {
			eventDescription = scanner.Text()
		}
		index++
	}
	err = bot.PutEvent(eventTime, eventDescription, username)
	if err != nil {
		bot.log.WithError(err).Error("Failed to add an event for user " + username)
		return err
	}
	return nil
}
