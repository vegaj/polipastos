package telegram

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gobuffalo/pop"
)

var config = make(map[string]string)

type ChatContext struct {
	ChatID int64 `json:"chat_id"`
	Offset int64 `json:"offset"`
}

//Conf allows to check the loaded configuration
func Conf(key string) string {
	return config[key]
}

//Init loads the bot configuration
func Init() {
	db, err := pop.Connect("configuration")
	if err != nil {
		panic(err.Error() + ". Check if the database is accessible.")
	}
	defer db.Close()

	var token string
	err = db.RawQuery("SELECT value FROM map WHERE name = ?", "telegram_bot_token").First(&token)
	if err != nil {
		panic(err)
	}

	config["bot_token"] = token
}

//Update loop for the bot
func Update() {
	api, err := tgbotapi.NewBotAPI(Conf("bot_token"))
	if err != nil {
		panic(err)
	}

	var conf tgbotapi.UpdateConfig

	updates, err := api.GetUpdates(conf)
	if err != nil {
		panic(err)
	}

	for _, update := range updates {

		if msg := update.Message; msg != nil {
			newmsg := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
			msgSent, err := api.Send(newmsg)
			var who string
			who = "Unknown"
			if err != nil {
				log.Println(err)
			} else {

				if msgSent.Contact != nil {
					who = msgSent.Contact.FirstName
				}
				log.Printf("Sent message: %v, %s to %s", msgSent.Date, msgSent.Text, who)
			}
		}
	}

}
