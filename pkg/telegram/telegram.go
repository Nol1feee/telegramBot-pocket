package telegram

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot{
	return &Bot{bot:bot}
}

func (b *Bot) Start() error {
	updates, err := b.initUpdatesChan()
	if err != nil {
		return err
	}

	b.handlerUpdates(updates)
	return nil
}

func (b *Bot) handlerUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {

		if update.Message == nil { // ignore any non message update 
			continue 
		}

		if update.Message.IsCommand() {
			b.handlerCommand(update)
			continue
		}

		b.handlerMessage(update)
	}
}

func (b *Bot) handlerCommand(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "i don't know this command")
	switch update.Message.Command(){
	case "start":
		msg.Text = "hi, start"
	default:
	}
	b.bot.Send(msg)
}

func (b *Bot) handlerMessage(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	b.bot.Send(msg)
}

func (b *Bot) initUpdatesChan() (tgbotapi.UpdatesChannel, error) {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	return updates, err
}