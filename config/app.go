package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	App struct {
		Name    string
		Version string
		Logger  *logrus.Logger
	}

	Telegram struct {
		Token string
	}
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading config file: %s", err)
	}

	App.Name = "Elwin's App"
	App.Version = "0.0.1"
	App.Logger = logrus.New()
	App.Logger.SetLevel(logrus.TraceLevel)

	Telegram.Token = os.Getenv("TELEGRAM_KEY")
}
