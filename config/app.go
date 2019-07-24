package config

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"

	_ "github.com/jessevdk/go-flags"
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
		Token       string
		WebhookHost string
	}
)

var opts struct {
	Env string `short:"e" long:"env" description:"Set path of .env"`
}

func init() {

	_, err := flags.Parse(&opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if opts.Env == "" {
		err = godotenv.Load()
	} else {
		err = godotenv.Load(opts.Env)
	}

	if err != nil {
		logrus.Fatal("Error loading config file: ", err)
	}

	App.Name = "Elwin's App"
	App.Version = "0.0.1"
	App.Logger = logrus.New()
	App.Logger.SetLevel(logrus.TraceLevel)
	App.Url = os.Getenv("APP_URL")

	Telegram.Token = os.Getenv("TELEGRAM_KEY")
	Telegram.WebhookHost = os.Getenv("TELEGRAM_URL")
}
