package redis

import (
	"github.com/go-redis/redis"
)

type Redis struct {
	Client *redis.Client
}

func NewRedisClient() Redis {
	return Redis{redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})}
}

func (redis *Redis) GetValue(k string) (string, bool) {
	result, ok := redis.Client.Get(k).Result()
	if ok != nil {
		return "", false
	}
	return result, true
}

func (redis *Redis) SetValue(k, v string) error {
	redis.Client.Set(k, v, 0)
	return nil
}
func (redis *Redis) StoreMap(k string, m map[string]interface{}) {
	redis.Client.HMSet(k, m)
}
func (redis *Redis) IsThere(url string) (string, bool) {

	return "0", false
}
