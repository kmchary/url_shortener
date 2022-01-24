package ports

import "context"

type CachePort interface {
	Set(ctx context.Context, key, value string) error
	SetIfNotExists(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) string
}
