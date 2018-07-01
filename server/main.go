package main

import (
	"time"

	"github.com/polipastos/server/telegram"
)

func main() {
	d, _ := time.ParseDuration("1s")
	telegram.Init()
	for i := 0; i < 3; i++ {
		telegram.Update()
		time.Sleep(d)
	}
}
