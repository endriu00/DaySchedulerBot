package main

import (
	"fmt"
	botservice "github.com/endriu00/DaySchedulerBot/service"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := run()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
}

func run() error {
	//Load configuration
	cfg, err := loadConfiguration()
	if err != nil {
		return err
	}

	//Start logger
	log := logrus.NewEntry(logrus.StandardLogger())

	//Connect to DB
	db, err := sqlx.Open("mysql", cfg.Database.DB)
	if err != nil {
		log.WithError(err).Error("Could not connect to DB")
		return err
	}

	//Make a channel to listen for an interrupt or terminate signal from the OS.
	//Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	//Make a channel to listen for errors coming from the listener. Use a
	//buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 2)

	//Initialize bot
	bot := botservice.New(&botservice.Config{
		DB:               db,
		Logger:           log,
		TelegramApiUrl:   cfg.Telegram.TelegramApiUrl,
		TelegramBotToken: cfg.Telegram.TelegramBotToken,
	})
	go func() {
		log.Info("Worker starting")
		serverErrors <- bot.ListenAndServe()
		log.Info("stopping API server")
	}()

	//Shutdown

	//Blocking main and waiting for shutdown signal or POSIX signals
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		log.Infof("signal %v received, start shutdown", sig)

		err := bot.Close()
		if err != nil {
			log.WithError(err).Warning("graceful shutdown of worker error")
		}

		// Log the status of this shutdown.
		switch {
		case sig == syscall.SIGSTOP:
			return errors.New("integrity issue caused shutdown")
		case err != nil:
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}
