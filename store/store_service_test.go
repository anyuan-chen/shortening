package store

import (
	"math/rand"
	"testing"
	"github.com/stretchr/testify/assert"
)
var testStore = &Redis{}

func init() {
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
	retrieved, _ := RetrieveUrl(shortURL)
	assert.Equal(t, longURL, retrieved)
}