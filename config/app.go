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
		Url     string
		Logger  *logrus.Logger
	}

	Telegram struct {
		Token string
	}
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading config file: ", err)
	}

	App.Name = "Elwin's App"
	App.Version = "0.0.1"
	App.Logger = logrus.New()
	App.Logger.SetLevel(logrus.TraceLevel)
	App.Url = os.Getenv("APP_URL")

	Telegram.Token = os.Getenv("TELEGRAM_KEY")
}
