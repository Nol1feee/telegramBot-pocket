package telegram

import (
	"fmt"
	"github.com/Nol1feee/telegramBot-pocket/internal/storage"
	"github.com/sirupsen/logrus"
	"log"

	// "github.com/Nol1feee/go-pocket-sdk/go-pocket-sdk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	errLinkMsg = "sorry, but your link is invalid"
	addLinkMsg = "add link to your pocket account!"
)

func (b *Bot) handlerUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {

		if update.Message == nil { // ignore any non message update
			continue
		}

		if update.Message.IsCommand() {
			if err := b.handlerCommand(update.Message); err != nil {
				log.Fatal(err)
			}
			continue
		}

		b.handlerMessage(update)
	}
}

func (b *Bot) handlerCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case "start":
		return b.handlerCommandStart(message)
	default:
		return b.handlerUnknownCommand(message)
	}
}

func (b *Bot) handlerCommandStart(message *tgbotapi.Message) error {
	authLink, err := b.generateAutozitaionLink(message.From.ID)
	logrus.Info(message.From.ID, "required authLink", authLink)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(startReply, authLink))

	b.bot.Send(msg)

	return err
}

func (b *Bot) handlerUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды :/")

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handlerMessage(update tgbotapi.Update) {
	token, err := b.storage.Get(update.Message.MessageID, storage.AccessToken)
	if err != nil {
		logrus.Errorf("access token error | %s", err)
		return
	}

	b.pocket.AccessToken = token
	err = b.pocket.Add(update.Message.Text)

	if err != nil {
		b.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, errLinkMsg))
		logrus.Info(err, "| can't add link, because than not a link")
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, addLinkMsg)

	b.bot.Send(msg)
}
