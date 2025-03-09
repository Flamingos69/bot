package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		default:
			defaultAnswer(bot, update.Message)
		}

	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help -> help")

	bot.Send(msg)
}

func defaultAnswer(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	if inputMessage != nil { // If we got a message
		log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
		// msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
