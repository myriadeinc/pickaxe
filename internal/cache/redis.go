package cache

import (
	"context"

	"github.com/spf13/viper"

	"github.com/go-redis/redis/v8"
)

type CacheService interface {
	SaveNewTemplate(map[string]interface{}) error
}

type RedisService struct {
	client *redis.Client
}

func NewClient() CacheService {

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &RedisService{
		client: rdb,
	}

}

func (r *RedisService) SaveNewTemplate(template map[string]interface{}) error {
	key := "blocktemplate"
	return r.client.Set(context.Background(), key, template, 0).Err()
}
