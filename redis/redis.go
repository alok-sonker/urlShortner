package redis

import (
	"fmt"
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
	fmt.Println("ok ", ok)
	if ok != nil {
		return "", false
	}
	fmt.Println("value stored is ", result)
	return result, true
}

func (redis *Redis) SetValue(k, v string) error {
	redis.Client.Set(k, v, 0)
	return nil
}
