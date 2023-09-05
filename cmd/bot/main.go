package main

import (
	"log"

	"github.com/Nol1feee/go-pocket-sdk/go-pocket-sdk"
	"github.com/Nol1feee/telegramBot-pocket/pkg/storage/boltdb"
	"github.com/Nol1feee/telegramBot-pocket/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5955951750:AAF2EfnxuNvNwOdtBEiRxI7rbgzhPTJ0JmY")
	if err != nil {
		log.Fatal(err)
	}

	pocket, err := pocket.NewClient("107399-dd37d8729062a0be4df6392", "https://t.me/apiPocketBot")
	if err != nil {
		log.Fatal(err)
	}

	//decode response and return slice of bytes 
	bot.Debug = true

	db, err := boltdb.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	TokenStorage := boltdb.NewTokenStorage(db)

	newBot := telegram.NewBot(bot, pocket, TokenStorage)
	err = newBot.Start(); if err != nil {
		log.Fatal(err)
	}
}