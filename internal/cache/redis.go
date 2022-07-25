package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"

	"github.com/go-redis/redis/v8"
)

type CacheService interface {
	SaveNewTemplate(map[string]interface{}) error
}

type StrictTemplate struct {
	BlockTemplateBlob string  `json:"blocktemplate_blob"`
	Difficulty        float64 `json:"difficulty"`
	SeedHash          string  `json:"seed_hash"`
	Height            float64 `json:"height"`
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

	redisTemplate := toRedisCompatibleStruct(template)
	return r.client.Set(context.Background(), key, redisTemplate, 0).Err()
}

func toRedisCompatibleStruct(template map[string]interface{}) StrictTemplate {
	return StrictTemplate{
		BlockTemplateBlob: fmt.Sprintf("%v", template["blocktemplate_blob"]),
		Difficulty:        template["difficulty"].(float64),
		SeedHash:          fmt.Sprintf("%v", template["seed_hash"]),
		Height:            template["height"].(float64),
	}
}

func (s StrictTemplate) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}
