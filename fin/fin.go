package fin

import (
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Fin(w http.ResponseWriter, r *http.Request) {
	defer w.WriteHeader(200)

	token := os.Getenv("TOKEN")

	bot, _ := tgbotapi.NewBotAPI(token)

	bot.Debug = false

	update, _ := bot.HandleUpdate(r)

	if !update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		bot.Send(msg)
	}
}
