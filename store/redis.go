package store

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(addr string) *RedisStore {
	return &RedisStore{client: redis.NewClient(&redis.Options{Addr: addr})}
}

func (r *RedisStore) Save(token, userID string, ttl time.Duration) error {
	return r.client.Set(context.Background(), token, userID, ttl).Err()
}

func (r *RedisStore) Get(token string) (string, bool) {
	uid, err := r.client.Get(context.Background(), token).Result()
	return uid, err == nil
}

func (r *RedisStore) Delete(token string) error {
	return r.client.Del(context.Background(), token).Err()
}
