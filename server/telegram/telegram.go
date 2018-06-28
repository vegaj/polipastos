package telegram

import (
	"github.com/gobuffalo/pop"
)

var config = make(map[string]string)

//Conf allows to check the loaded configuration
func Conf(key string) string {
	return config[key]
}

//Init loads the bot configuration
func Init() {
	db, err := pop.Connect("configuration")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var token string
	err = db.RawQuery("SELECT value FROM map WHERE name = ?", "telegram_bot_token").First(&token)
	if err != nil {
		panic(err)
	}

	config["bot_token"] = token
}
