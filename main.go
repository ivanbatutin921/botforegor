package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"root/commands"
	"root/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6753629557:AAFYqNxfYFLpAzPKtjLOo74703yg2bo6_3o")
	if err != nil {
		fmt.Println("1")
		log.Fatal(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		fmt.Println("2")
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Chat.ID != 0 {
			chatID := update.Message.Chat.ID
			os.Setenv("CHAT_ID", strconv.FormatInt(chatID, 10))
		}

		if update.Message.Text == "/start" {
			handlers.SayHello(bot)
			commands.SelectTemplate(bot)
		}

		if update.Message.Text == "/sendTwoPhoto" {
			handlers.SendTwoPhoto(bot)	
		}

	}
}
