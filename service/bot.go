package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"sync"
)

type Bot struct {
	db               *sqlx.DB
	log              *logrus.Entry
	telegramApiUrl   string
	telegramBotToken string

	shutdownSignal chan interface{}
	shutdown       bool
	wg             sync.WaitGroup
	wgchan         []chan interface{}
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

		shutdownSignal: make(chan interface{}, 1),
		shutdown:       false,
		wgchan:         make([]chan interface{}, 0),
	}
}

func (bot *Bot) Close() error {
	return nil
}

func (bot *Bot) ListenAndServe() error {
	bot.log.Info("Bot : Listen And Serve")
	var shutdown = make(chan interface{}, 1)
	bot.wg.Add(1)
	go func() {
		for !bot.shutdown {
			<-shutdown
		}
		bot.wg.Done()
	}()
	bot.wgchan = append(bot.wgchan, shutdown)

	//bot.HandleTelegramWebhook()
	// Wait for shutdown signal
	<-bot.shutdownSignal
	bot.log.Info("Worker : Stopping scheduled tasks")
	bot.shutdown = true
	for _, c := range bot.wgchan {
		c <- 0
	}
	bot.log.Info("Worker : Waiting for tasks")
	bot.wg.Wait()
	return nil
}
