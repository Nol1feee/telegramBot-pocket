package boltdb

import (

	storage "github.com/Nol1feee/telegramBot-pocket/pkg/storage"
	"github.com/boltdb/bolt"
)

func InitDB() (*bolt.DB, error) {
	db, err := bolt.Open("token.db", 0600, nil)
	if err != nil {
		return &bolt.DB{}, err
	}
	// defer db.Close()

	if err := db.Batch(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(storage.AccessToken))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(storage.RequestToken))
		return err
	}); err != nil {
		return nil, err
	}

	return db, nil
}