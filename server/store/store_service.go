package store

import (
	"context"
	"fmt"
	"os"
	"time"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	redisClient *redis.Client
}
var (
	storeService = &Redis{}
	ctx = context.Background()
)
const CacheDuration = 6 * time.Hour

func InitializeStore() *Redis {
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: 0,
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error initializing Redis: %v", err))
	}
	storeService.redisClient = redisClient
	return storeService
}

func InsertUrl(short string, original string, userId string){
	err := storeService.redisClient.Set(ctx, short, original, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed to save a specific redis key: error: %v - shortUrl: %v - originalUrl: %v", err, short, original))
	}
}

func RetrieveUrl (short string) string{
	result, err := storeService.redisClient.Get(ctx, short).Result()
	if err != nil {
		if err == redis.Nil {
			panic(fmt.Sprintf("this key doesn't exist in redis"))
		}
		panic(fmt.Sprintf("RetriveUrl fail - Error: %v - ShortUrl: %v", err, short))
	}
	return result
}
