package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TelegramBotConfiguration struct {
	Config struct {
		Path string
	}
	Database struct {
		DB *sqlx.DB
	}

	Logger *logrus.Entry
}
