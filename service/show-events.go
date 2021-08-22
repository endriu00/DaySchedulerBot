package service

import (
	"net/http"
	"net/url"
	"strconv"
)

/*
  ShowEvents separates user request (identified by 'chat' chatId) having 'day'
  as the day from which fetch data.
*/
func (bot *Bot) ShowEvents(chat Chat, day string) error {
	var err error
	var events []Event
	if day == "today" {
		events, err = bot.GetEventsToday(chat)
		if err != nil {
			bot.log.WithError(err).Error("Could not fetch today events.")
			return err
		}
	}
	if day == "tomorrow" {
		events, err = bot.GetEventsTomorrow(chat)
		if err != nil {
			bot.log.WithError(err).Error("Could not fetch tommorrow events.")
			return err
		}
	}
	//if day == "week" {
	//	ShowEventsWeek()
	//}

	//Create response message
	telegramMessage := bot.BuildMessage(events)

	telegramApi := bot.telegramApiUrl + bot.telegramBotToken + "/sendMessage"
	telegramApiURL, err := url.Parse(telegramApi)
	if err != nil {
		bot.log.WithError(err).Error("Failed to parse Telegram API url")
		return err
	}
	_, err = http.PostForm(
		telegramApiURL.String(),
		url.Values{
			"chat_id": {strconv.Itoa(chat.Id)},
			"text":    {telegramMessage},
		})
	if err != nil {
		bot.log.WithError(err).Error("Failed sending telegram message.")
		return nil
	}
	//Send message to user with the events fetched
	return nil
}
