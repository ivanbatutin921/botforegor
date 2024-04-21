package handlers

import (
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	photo1      *[]tgbotapi.PhotoSize
	photo2      *[]tgbotapi.PhotoSize
	secondPhoto bool
)

func SayHello(bot *tgbotapi.BotAPI) {
	msgChatID, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	msg := tgbotapi.NewMessage(msgChatID, "Привет, я бот для обработки шаблонов")
	bot.Send(msg)
}

func SendTwoPhoto(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	msgChatID, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)

	var secondPhoto bool

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Photo != nil && !secondPhoto {
			photo1 = update.Message.Photo
			msg := tgbotapi.NewMessage(msgChatID, "Отправьте вторую фотографию")
			bot.Send(msg)
			secondPhoto = true

		} else if update.Message.Photo != nil && secondPhoto {
			photo2 = update.Message.Photo

			msg := tgbotapi.NewMessage(msgChatID, "Обе фотографии загружены")
			bot.Send(msg)
			GetPicture(bot, photo1, photo2)
			GenerateStickerPack(bot)

		} else {
			msg := tgbotapi.NewMessage(msgChatID, "Просим отправить фото")
			bot.Send(msg)
			secondPhoto = false
		}
	}
}

func GetDataForStickerPack() {

}
