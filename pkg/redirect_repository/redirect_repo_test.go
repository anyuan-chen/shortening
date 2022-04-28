package redirectrepository_test

import (
	"os"
	"testing"

	"github.com/anyuan-chen/urlshortener/server/pkg/redirect_repository/redis"
	"github.com/stretchr/testify/assert"
)

var test redis.RedisRedirectRepository

func init(){
	var err error
	test, err = redis.CreateRedisRepository(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"))
	if err != nil {
		panic("failed to initialize redis")
	}
}
func TestGetAndCreateAndDelete(t *testing.T){
	shortened_link := "short"
	original_link := "long"
	test.Create(shortened_link, original_link, "guest")
	original_redis, err := test.Get(shortened_link)
	assert.Equal(t, original_link, original_redis)
	assert.Nil(t, err)
	test.Delete(shortened_link)
	_, err = test.Get(shortened_link)
	assert.NotNil(t, err)
}

