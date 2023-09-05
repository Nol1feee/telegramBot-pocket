package telegram

import (
	// "fmt"
	"log"

	"github.com/Nol1feee/go-pocket-sdk/go-pocket-sdk"
	"github.com/Nol1feee/telegramBot-pocket/pkg/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	startReply = "Утречка!\n\n Чтобы воспользоваться моим функционалом предоставь частичный доступ к pocket аккаунту, - перейди по следующей ссылке и подтверди свое согласие, спасибо!\n%s"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	pocket *pocket.Client
	storage storage.TokenStorage
}

func NewBot(bot *tgbotapi.BotAPI, pocket *pocket.Client, token storage.TokenStorage) *Bot{
	return &Bot{bot:bot, pocket:pocket, storage: token}
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
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	return updates, err
}