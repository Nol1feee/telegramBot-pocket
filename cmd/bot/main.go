package main

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/Nol1feee/pkg/telegram"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5955951750:AAF2EfnxuNvNwOdtBEiRxI7rbgzhPTJ0JmY")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	newBot := telegram.NewBot(bot)
	err = newBot.Start(); if err != nil {
		log.Fatal(err)
	}
}