package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// type RedirectRepository interface {
// 	Get(shortened_link string) (string, error)  //returns the longer link
// 	Create(shortened_link string, original_link string, user_id string) error
// }

type redisRedirectRepository struct {
	redis *redis.Client
}

func CreateRedisRepository(redis_address string, redis_password string) (redisRedirectRepository, error){
	redisClient := redis.NewClient(&redis.Options{
		Addr: redis_address,
		Password: redis_password,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return redisRedirectRepository{}, err;
	}
	redisRepo := redisRedirectRepository {
		redisClient,
	}
	return redisRepo ,err
}

func (r *redisRedirectRepository) Get(shortened_link string) (string, error){
	
}

func (r *redisRedirectRepository) Create(shortened_link string, original_link string, user_id string){

}