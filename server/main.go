package main

import (
	"log"

	"github.com/polipastos/server/telegram"
)

func main() {
	telegram.Init()
	log.Println("Bot token", telegram.Conf("bot_token"))
}
