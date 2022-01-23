package caching

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type CachingStore interface {
	Set(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
	SetNX(ctx context.Context, key, value string) error // set key, value if the key value does not already exist
}

func NewCashingStore(ctx context.Context, host, port string) CachingStore {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &redisStore{db: redisClient}
}

type redisStore struct {
	db *redis.Client
}

func (s *redisStore) Set(ctx context.Context, key, value string) error {
	return s.db.Set(ctx, key, value, 0).Err()
}

func (s *redisStore) SetNX(ctx context.Context, key, value string) error {
	return s.db.SetNX(ctx, key, value, 0).Err()
}

func (s *redisStore) Get(ctx context.Context, key string) (string, error) {
	val, err := s.db.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
