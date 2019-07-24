package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"

	"github.com/gin-gonic/autotls"

	"github.com/elwin/messenger/config"
	"github.com/gin-gonic/gin"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type App struct {
	logger logrus.FieldLogger
	router *gin.Engine
	bot    *telegram.BotAPI
}

var messages []string

func main() {

	app := &App{}

	app.logger = config.App.Logger
	app.router = SetupRouter(app)

	var updates telegram.UpdatesChannel
	app.bot, updates = SetupTelegram(app)

	go func() {
		for update := range updates {
			if update.Message == nil {
				app.logger.Info("Received nil message")
				continue
			}

			logger := app.logger.WithFields(logrus.Fields{
				"id":      update.Message.Chat.ID,
				"user":    update.Message.Chat.FirstName,
				"message": update.Message.Text,
			})

			logger.Info("Received message")

			messages = append(messages, update.Message.Text)

			text := "shut the fuck up, " + update.Message.Chat.FirstName + "!"

			if update.Message.Text == "p==np?" || update.Message.Text == "p=np?" {
				if rand.Intn(2) == 1 {
					text = "Wahrschinli nid"
				} else {
					text = "Auä scho"
				}
			}

			message := telegram.NewMessage(update.Message.Chat.ID, text)
			_, err := app.bot.Send(message)
			if err != nil {
				logger.Error("Failed to send reply")
			}
		}
	}()

	autotls.Run(app.router)
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": messages,
	})
}

func SetupRouter(app *App) *gin.Engine {
	r := gin.Default()

	r.GET("/", home)

	return r
}

func SetupTelegram(app *App) (*telegram.BotAPI, telegram.UpdatesChannel) {

	bot, err := telegram.NewBotAPI(config.Telegram.Token)
	if err != nil {
		app.logger.Fatal("Failed to initialize Telegram Bot with API Key")
	}

	url := config.Telegram.WebhookHost + "/webhook"

	ch := make(chan telegram.Update, bot.Buffer)

	app.router.POST("/webhook", func(c *gin.Context) {
		bytes, _ := ioutil.ReadAll(c.Request.Body)

		var update telegram.Update
		json.Unmarshal(bytes, &update)

		ch <- update
	})

	_, err = bot.SetWebhook(telegram.NewWebhook(url))
	if err != nil {
		app.logger.Error("Failed to set webhook: ", err)
	}

	return bot, ch

}
