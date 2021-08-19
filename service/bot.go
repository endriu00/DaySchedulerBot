package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	db               *sqlx.DB
	log              *logrus.Entry
	telegramApiUrl   string
	telegramBotToken string
}

type Config struct {
	DB               *sqlx.DB
	Logger           *logrus.Entry
	TelegramApiUrl   string
	TelegramBotToken string
}

func New(cfg *Config) Bot {
	return Bot{
		db:               cfg.DB,
		log:              cfg.Logger,
		telegramApiUrl:   cfg.TelegramApiUrl,
		telegramBotToken: cfg.TelegramBotToken,
	}
}
