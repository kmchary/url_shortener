package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Adapter struct {
	db *redis.Client
}

func NewAdapter(host, port string) *Adapter {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Adapter{db: redisClient}
}

func (s *Adapter) Set(ctx context.Context, key, value string) error {
	return s.db.Set(ctx, key, value, 0).Err()
}

func (s *Adapter) SetIfNotExists(ctx context.Context, key, value string) error {
	return s.db.SetNX(ctx, key, value, 0).Err()
}

func (s *Adapter) Get(ctx context.Context, key string) (string, error) {
	val, err := s.db.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
