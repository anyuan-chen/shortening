package store

import (
	"math/rand"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)
var testStore = &Redis{}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(fmt.Sprintln("big problems!"))
	}
	testStore = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.NotNil(t, testStore.redisClient)
}

func TestInsertionAndRetrieval(t *testing.T) {
	randomString := func (length int) string {
		const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		b := make([]byte, length)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		return string(b)
	} 
	shortURL := randomString(8)
	longURL := randomString(8)
	userUUId := "andrew"
	InsertUrl(shortURL, longURL, userUUId)
	retrieved := RetrieveUrl(shortURL)
	assert.Equal(t, longURL, retrieved)
}