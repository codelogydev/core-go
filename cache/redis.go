package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Init(url string) error {
	opt, err := redis.ParseURL(url)
	if err != nil {
		return err
	}

	Client = redis.NewClient(opt)

	return Client.Ping(context.Background()).Err()
}

func Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return Client.Set(ctx, key, value, ttl).Err()
}

func Get(ctx context.Context, key string) (string, error) {
	return Client.Get(ctx, key).Result()
}

func Delete(ctx context.Context, keys ...string) error {
	return Client.Del(ctx, keys...).Err()
}

func Exists(ctx context.Context, key string) (bool, error) {
	n, err := Client.Exists(ctx, key).Result()
	return n > 0, err
}
