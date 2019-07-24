package main

import (
	"math/rand"
	"net/http"

	"github.com/elwin/messenger/config"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func main() {

	bot, err := telegram.NewBotAPI(config.Telegram.Token)
	logger := config.App.Logger

	if err != nil {
		logger.Fatal("Failed to initialize Telegram Bot with API Key")
	}

	url := "https://659e7378.ngrok.io/webhook"

	_, err = bot.SetWebhook(telegram.NewWebhook(url))
	if err != nil {
		logger.Error("Failed to set webhook: %s", err)
	}

	go http.ListenAndServe("localhost:4444", nil)
	logrus.Info("Listening on localhost:4444")

	updates := bot.ListenForWebhook("/webhook")

	for update := range updates {
		if update.Message == nil {
			logger.Info("Received nil message")
			continue
		}

		text := "shut the fuck up, " + update.Message.Chat.FirstName + "!"

		if update.Message.Text == "p==np?" || update.Message.Text == "p=np?" {
			if rand.Intn(2) == 1 {
				text = "Wahrschinli nid"
			} else {
				text = "Auä scho"
			}
		}

		message := telegram.NewMessage(update.Message.Chat.ID, text)
		_, err := bot.Send(message)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"id":   message.ChatID,
				"user": message.ChannelUsername,
			}).Error("Failed to send reply")
		}
	}
}
