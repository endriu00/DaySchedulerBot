package service

/*
Update type is the type mirroring the concept of the request
Telegram servers send to the webhook.
It has two elements in it:
	- UpdateId that is a unique id identifying the update message
	- Message that is an object of Message type; the message contained in the request
*/
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

/*
Message type is the type representing the effective message sent
by the user talking to the bot.
It has two elements:
	- Text that is a string containing the text sent by the user
	- Chat that is the chat object in which the Message is
*/
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

/*
Chat type represents a chat between the user and the bot.
In other words, it identifies the user the bot is talking to.
	- Id is the id of the chat
*/
type Chat struct {
	Id int `json:"id"`
}
