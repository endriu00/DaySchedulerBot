package service

import (
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func (bot *Bot) GetEventsToday(chat Chat) ([]Event, error) {
	id := chat.Id
	var events []Event
	var err error
	//Query the DB for user events of today
	err = bot.db.Select(&events, `SELECT username, time, description 
			  FROM event JOIN agenda ON event.username=agenda.username 
			  AND event.time=agenda.time 
			  WHERE event.time=CURDATE() AND agenda.chatid=?`, id)
	//TODO: err can be many things. Distinguish them!
	if err != nil {
		bot.log.WithError(err).Error("Failed querying for today events of user " + strconv.Itoa(id))
		return nil, err
	}
	return events, nil
}
