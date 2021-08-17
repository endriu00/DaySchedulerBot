package service

import ()

/*
  ShowEvents separates user request (identified by 'chat' chatId) having 'day'
  as the day from which fetch data.
*/
func ShowEvents(chat Chat, day string) error {
	var err error
	if day == "today" {
		err = GetEventsToday(chat)
		if err != nil {
			return err
		}
	}
	if day == "tomorrow" {
		err = GetEventsTomorrow(chat)
		if err != nil {
			return err
		}
	}
	//if day == "week" {
	//	ShowEventsWeek()
	//}
	//Query db for events of the chatId user

	//Send message to user with the events fetched
	return nil
}
