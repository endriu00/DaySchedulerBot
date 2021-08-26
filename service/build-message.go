package service

import ()

func (bot *Bot) BuildMessage(events []Event) string {
	var message string
	username := events[0].Username

	message += "Dear @" + username + " let's look at your events."
	for _, event := range events {
		//TODO: improve memory management of string creation. Use string builder for example, or strings join method
		message += event.TimeScheduled.String() + " " + event.Description + "\n"
	}
	return message
}
