package telegram

import (
	"fmt"
	"github.com/Nol1feee/telegramBot-pocket/internal/storage"
)

func (b *Bot) generateAutozitaionLink(userId int) (string, error) {
	reqToken, err := b.pocket.GetRequestToken()
	if err != nil {
		return "", err
	}

	if err := b.storage.Save(userId, reqToken, storage.RequestToken); err != nil {
		return "", err
	}

	//redirectUrl := fmt.Sprintf("localhost:80/auth?user_id=%d", userId)
	redirectUrl := fmt.Sprintf("google.com")
	return b.pocket.GetAutorizationUrl(reqToken, redirectUrl)
}
