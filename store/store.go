package store

import "time"

type TokenStore interface {
	Save(token string, userID string, ttl time.Duration) error
	Get(token string) (string, bool)
	Delete(token string) error
}
