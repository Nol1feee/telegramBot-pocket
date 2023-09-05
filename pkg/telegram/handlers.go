package telegram

import (
	"fmt"
	"log"

	// "github.com/Nol1feee/go-pocket-sdk/go-pocket-sdk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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
	switch message.Command(){
	case "start":
		return b.handlerCommandStart(message)
	default:
		return b.handlerUnknownCommand(message)
	}
}

func (b *Bot) handlerCommandStart(message *tgbotapi.Message) error {
	authLink, err := b.generateAutozitaionLink(message.MessageID)
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
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	b.bot.Send(msg)
}