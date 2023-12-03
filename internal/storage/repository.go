package storage

type Bucket string

const (
	AccessToken  Bucket = "access_token"
	RequestToken Bucket = "request_token"
)

type TokenStorage interface {
	Save(userId int, token string, bucket Bucket) error
	Get(userId int, bucket Bucket) (string, error)
}
