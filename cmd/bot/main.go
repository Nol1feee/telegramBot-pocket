package main

import (
	"github.com/Nol1feee/telegramBot-pocket/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal("can't run")
	}
}
