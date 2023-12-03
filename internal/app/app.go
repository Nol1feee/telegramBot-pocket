package app

import (
	pocket "github.com/Nol1feee/telegramBot-pocket/internal/api/pocketSDK"
	"github.com/Nol1feee/telegramBot-pocket/internal/authServer"
	boltdb2 "github.com/Nol1feee/telegramBot-pocket/internal/storage/boltdb"
	"github.com/Nol1feee/telegramBot-pocket/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"os"
)

const redirectUri = "localhost:80/auth?user_id=%s"

func Run() error {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TGBOT_TOKEN"))
	if err != nil {
		return err
	}

	pocket, err := pocket.NewClient(os.Getenv("CONSURMER_KEY"), redirectUri)
	if err != nil {
		return err
	}

	//decode response and return slice of bytes
	//bot.Debug = true

	db, err := boltdb2.InitDB()
	if err != nil {
		return err
	}

	TokenStorage := boltdb2.NewTokenStorage(db)

	srv := authServer.NewServer(pocket, TokenStorage)
	go srv.Start()
	logrus.Info("auth server working")

	newBot := telegram.NewBot(bot, pocket, TokenStorage, srv)
	err = newBot.Start()
	if err != nil {
		return err
	}

	return nil
}
