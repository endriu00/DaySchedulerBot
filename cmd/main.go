package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"os"
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

	//Connect to DB
	db, err := sqlx.Open("mysql", "connectionstringtobeadded") //add connection string
	if err != nil {
		return err
	}

	//Set specified configurations to the bot
	return nil
}
