package service

func Handler(u *Update) error {
	msg := u.Message.Text
	command := SanityzeCommand(msg)

}
