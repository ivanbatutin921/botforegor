package commands

import (
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SelectTemplate(bot *tgbotapi.BotAPI) {

	templates := []string{"template1", "template2", "template3"}

	//keyboard := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonWithRequestContact("Выберите шаблон"))

	keyboard := tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{
		{Text: "Выберите шаблон"},
	})

	kbButtons := make([]tgbotapi.KeyboardButton, 0, len(templates))
	for _, template := range templates {
		kbButtons = append(kbButtons, tgbotapi.KeyboardButton{Text: template, RequestContact: true, RequestLocation: true})
	}
	keyboard.Keyboard = append(keyboard.Keyboard, []tgbotapi.KeyboardButton{kbButtons[0], kbButtons[1], kbButtons[2]})



	msgChatID, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)

	msg1 := tgbotapi.NewMessage(msgChatID, "Выберите шаблон для обработки")
	msg1.ReplyMarkup = keyboard
	bot.Send(msg1)
}
