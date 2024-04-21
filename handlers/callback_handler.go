package handlers

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func GetPicture(bot *tgbotapi.BotAPI, photo1, photo2 *[]tgbotapi.PhotoSize) {
	// Скачиваем первый изображение
	file1, err := bot.GetFile(tgbotapi.FileConfig{FileID: (*photo1)[len(*photo1)-1].FileID})
	if err != nil {
		log.Println("Ошибка: не удалось скачать файл. \n", err)
		return
	}
	// Скачиваем второй изображение
	file2, err := bot.GetFile(tgbotapi.FileConfig{FileID: (*photo2)[len(*photo2)-1].FileID})
	if err != nil {
		log.Println("Ошибка: не удалось скачать файл. \n", err)
		return
	}

	// Создаем POST-запрос на сервер
	req, err := http.NewRequest("POST", "http://localhost:5000/swap", nil)
	if err != nil {
		log.Println("Ошибка: не удалось создать запрос. \n", err)
		return
	}

	// Создаем форму для запроса
	form := &bytes.Buffer{}
	writer := multipart.NewWriter(form)

	// Добавляем первый изображение в форму
	fw1, err := writer.CreateFormFile("image1", "image1.jpg")
	if err != nil {
		log.Println("Ошибка: не удалось создать поле формы. \n", err)
		return
	}
	// Скачиваем первый изображение с сервера Telegram
	resp, err := http.Get("https://api.telegram.org/file/bot" + bot.Token + "/" + file1.FilePath)
	if err != nil {
		log.Println("Ошибка: не удалось скачать файл. \n", err)
		return
	}
	defer resp.Body.Close()
	// Копируем содержимое файла в поле формы
	_, err = io.Copy(fw1, resp.Body)
	if err != nil {
		log.Println("Ошибка: не удалось копировать содержимое файлов. \n", err)
		return
	}
	// Добавляем второй изображение в форму
	fw2, err := writer.CreateFormFile("image2", "image2.jpg")
	if err != nil {
		log.Println("Ошибка: не удалось создать поле формы. \n", err)
		return
	}
	// Скачиваем второй изображение с сервера Telegram
	resp, err = http.Get("https://api.telegram.org/file/bot" + bot.Token + "/" + file2.FilePath)
	if err != nil {
		log.Println("Ошибка: не удалось скачать файл. \n", err)
		return
	}
	defer resp.Body.Close()
	// Копируем содержимое файла в поле формы
	_, err = io.Copy(fw2, resp.Body)
	if err != nil {
		log.Println("Ошибка: не удалось копировать содержимое файлов. \n", err)
		return
	}
	// Закрываем запись в форму
	writer.Close()

	// Устанавливаем заголовок запроса
	req.Body = io.NopCloser(form)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Отправляем запрос на сервер
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Println("Ошибка: не удалось отправить запрос. \n", err)
		return
	}
	defer resp.Body.Close()

	// Проверяем статус код ответа
	if resp.StatusCode == http.StatusOK {
		// Создаем папку uploads, если она еще не существует
		err := os.MkdirAll("uploads", 0755)
		if err != nil {
			log.Println("Ошибка: не удалось создать папку uploads. \n", err)
			return
		}

		// Сохраняем ответ сервера в файл
		file, err := os.Create("uploads/image.png")
		if err != nil {
			log.Println("Ошибка: не удалось создать файл. \n", err)
			return
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Println("Ошибка: не удалось сохранить файл. \n", err)
			return
		}

		log.Println("Файл успешно сохранен в папку uploads.")
	}

}

func GenerateStickerPack(bot *tgbotapi.BotAPI) {
	//stickers := make([]tgbotapi.Sticker, 0)

}
