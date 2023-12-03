package boltdb

import (
	"github.com/Nol1feee/telegramBot-pocket/internal/storage"
	"strconv"

	"errors"
	"github.com/boltdb/bolt"
)

type TokenStorage struct {
	db *bolt.DB
}

func NewTokenStorage(db *bolt.DB) *TokenStorage {
	return &TokenStorage{db: db}
}

func (s *TokenStorage) Save(chatID int, token string, bucket storage.Bucket) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToBytes(chatID), []byte(token))
	})
}

func (t *TokenStorage) Get(userId int, bucket storage.Bucket) (string, error) {
	var token string

	err := t.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(storage.RequestToken))
		token = string(b.Get(intToBytes(userId)))
		return nil
	})

	if err != nil {
		return "", err
	}

	if token == "" {
		return "", errors.New("token by userId not found")
	}

	return token, nil

}

func intToBytes(chatId int) []byte {
	return []byte(strconv.Itoa(chatId))
}
