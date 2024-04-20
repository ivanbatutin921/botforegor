package handlers

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	photo1 *[]tgbotapi.PhotoSize
	photo2 *[]tgbotapi.PhotoSize
)

func SayHello(bot *tgbotapi.BotAPI) {
	msgChatID, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	msg := tgbotapi.NewMessage(msgChatID, "Привет, я бот для обработки шаблонов")
	bot.Send(msg)
}

func SendTwoPhoto(bot *tgbotapi.BotAPI) {
	msgChatID, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	var secondPhoto bool

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Photo != nil && !secondPhoto {
			photo1 = update.Message.Photo
			msg := tgbotapi.NewMessage(msgChatID, "Успешно1")
			bot.Send(msg)
			secondPhoto = true

		} else if update.Message.Photo != nil && secondPhoto {
			photo2 = update.Message.Photo

			msg := tgbotapi.NewMessage(msgChatID, "Успешно2")
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(msgChatID, "Просим отправить фото")
			bot.Send(msg)
			secondPhoto = false
		}

	}

}

// package main

// import (
// 	"log"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
// )

// func main() {
// 	bot, err := tgbotapi.NewBotAPI("6753629557:AAFYqNxfYFLpAzPKtjLOo74703yg2bo6_3o")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	bot.Debug = true
// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60
// 	updates, err := bot.GetUpdatesChan(u)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for update := range updates {
// 		if update.Message == nil {
// 			continue
// 		}
// 		if update.Message.Photo != nil {
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Успешно")

// 			bot.Send(msg)
// 		} else {
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Просим отправить фото")
// 			bot.Send(msg)
// 		}
// 	}
// }
