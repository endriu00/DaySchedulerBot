package main

import (
	"fmt"
	"github.com/ardanlabs/conf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type TelegramBotConfiguration struct {
	Config struct {
		Path string `conf:"default:/conf/config.yml"`
	}
	Database struct {
		DB string
	}
	Telegram struct {
		TelegramApiUrl   string `conf:"default:https://api.telegram.org/bot"`
		TelegramBotToken string
	}

	Logger *logrus.Entry
}

func loadConfiguration() (TelegramBotConfiguration, error) {
	// Create configuration defaults
	var cfg TelegramBotConfiguration

	// Try to load configuration from environment variables and command line switches
	if err := conf.Parse(os.Args[1:], "CFG", &cfg); err != nil {
		if err == conf.ErrHelpWanted {
			usage, err := conf.Usage("CFG", &cfg)
			if err != nil {
				return cfg, errors.Wrap(err, "generating config usage")
			}
			fmt.Println(usage)
			return cfg, conf.ErrHelpWanted
		}
		return cfg, errors.Wrap(err, "parsing config")
	}

	// Override values from YAML if specified and if it exists (useful in k8s/compose)
	fp, err := os.Open(cfg.Config.Path)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Can't read the config file, while it exists...")
		return cfg, err
	} else if err == nil {
		yamlFile, err := ioutil.ReadAll(fp)
		if err != nil {
			fmt.Printf("can't read config file: %v", err)
			return cfg, err
		}
		err = yaml.Unmarshal(yamlFile, &cfg)
		if err != nil {
			fmt.Printf("can't unmarshal config file: %v", err)
			return cfg, err
		}
		_ = fp.Close()
	}

	return cfg, nil
}
