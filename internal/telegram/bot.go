package telegram

import (
	"github.com/Nol1feee/telegramBot-pocket/internal/authServer"
	"github.com/Nol1feee/telegramBot-pocket/internal/storage"
	"github.com/sirupsen/logrus"

	pocket "github.com/Nol1feee/telegramBot-pocket/internal/api/pocketSDK"
	//"github.com/Nol1feee/go-pocket-sdk/go-pocket-sdk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	startReply = "Утречка!\n\n Чтобы воспользоваться моим функционалом предоставь частичный доступ к pocket аккаунту, - перейди по следующей ссылке и подтверди свое согласие, спасибо!\n%s"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	pocket  *pocket.Client
	storage storage.TokenStorage
	server  *authServer.AuthServer
}

func NewBot(bot *tgbotapi.BotAPI, pocket *pocket.Client, token storage.TokenStorage, authServer *authServer.AuthServer) *Bot {
	return &Bot{bot: bot, pocket: pocket, storage: token, server: authServer}
}

func (b *Bot) Start() error {
	updates, err := b.initUpdatesChan()
	if err != nil {
		return err
	}

	b.handlerUpdates(updates)
	return nil
}

func (b *Bot) initUpdatesChan() (tgbotapi.UpdatesChannel, error) {
	logrus.Info("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	return updates, err
}
