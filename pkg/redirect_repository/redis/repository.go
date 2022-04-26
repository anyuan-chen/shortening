package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// type RedirectRepository interface {
// 	Get(shortened_link string) (string, error)  //returns the longer link
// 	Create(shortened_link string, original_link string, user_id string) error
// }

type RedisRedirectRepository struct {
	redis *redis.Client
}

func CreateRedisRepository(redis_address string, redis_password string) (RedisRedirectRepository, error){
	redisClient := redis.NewClient(&redis.Options{
		Addr: redis_address,
		Password: redis_password,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return RedisRedirectRepository{}, err;
	}
	redisRepo := RedisRedirectRepository {
		redisClient,
	}
	return redisRepo ,err
}

func (r *RedisRedirectRepository) Get(shortened_link string) (string, error){
	original_link, err := r.redis.Get(context.Background(), shortened_link).Result()
	if err != nil {
		return "", err
	}
	return original_link, nil
}

func (r *RedisRedirectRepository) Create(shortened_link string, original_link string, user_id string) error{
	err := r.redis.Set(context.Background(), shortened_link, original_link, 0)
	if err != nil {
		msg := fmt.Sprintf("failed to insert shortened %s and original %s into redis", shortened_link, original_link) 
		return errors.New(msg)
	}
	return nil
}
// type SessionRepository interface {
// 	CreateSessionRepository()
// 	GetSession(session_id string) (Session, error)
// 	GetId(session_id string) (string, error)
// 	IsLoggedIn(session_id string) (Link, error)
// 	CreateSession(access_token string, refresh_token string, token_type string, expiry string, provider string) (string, error) 
// 	GetLoginRedirect(provider string, oauthstate string) (string, error)
// 	CodeExchange(provider string, code string) ([]byte, error)
// }