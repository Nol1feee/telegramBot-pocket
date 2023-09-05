package telegram

import "github.com/Nol1feee/telegramBot-pocket/pkg/storage"

func (b *Bot) generateAutozitaionLink(chatID int) (string, error) {
	reqToken, err := b.pocket.GetRequestToken()
	if err != nil {
		return "", err
	}

	if err := b.storage.Save(chatID, reqToken, storage.RequestToken); err != nil {
		return "", err
	}

	return b.pocket.GetAutorizationUrl(reqToken)
}