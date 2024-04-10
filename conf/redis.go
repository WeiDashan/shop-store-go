package conf

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

var rdClient *redis.Client
var monthDuration = 30 * 24 * 60 * 60 * time.Second

type RedisClient struct {
}

func InitRedis() (*RedisClient, error) {
	rdClient = redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		DB:   viper.GetInt("redis.database"),
	})

	_, err := rdClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{}, nil
}

func (rc *RedisClient) Set(key string, value any, rest ...any) error {
	duration := monthDuration
	if len(rest) > 0 {
		if v, ok := rest[0].(time.Duration); ok {
			duration = v
		}
	}
	return rdClient.Set(context.Background(), key, value, duration).Err()
}
func (rc *RedisClient) Get(key string) (any, error) {
	return rdClient.Get(context.Background(), key).Result()
}
func (rc *RedisClient) Delete(key ...string) error {
	return rdClient.Del(context.Background(), key...).Err()
}
