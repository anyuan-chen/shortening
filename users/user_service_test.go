package users

import (
	"log"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)
var testStore = &Database{}

func init() {
	testStore = InitializeDatabase()
}

func TestStoreInit(t *testing.T) {
	assert.NotNil(t, testStore.db)
}

func TestGetUser(t *testing.T) {
	id := "guest"
	retid, pfp := GetUser(id)
	assert.Equal(t, id, retid)
	assert.Equal(t, pfp, "")
}

func TestCreateUser(t *testing.T) {
	randomString := func (length int) string {
		const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		b := make([]byte, length)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		return string(b)
	} 
	id := randomString(8);
	err := CreateUser(id, "");
	if err != nil {
		log.Fatal(err)
	}
	retid, pfp := GetUser(id)
	assert.Equal(t, retid, id)
	assert.Equal(t, pfp, "")
	deleteUser(id)
}


func TestAddAndGetLinksByUser(t *testing.T) {
	randomString := func (length int) string {
		const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		b := make([]byte, length)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		return string(b)
	} 
	id := randomString(8)
	err := CreateUser(id, "")
	if err != nil {
		log.Fatal(err)
	}
	AddLink(id, "one", "two", "his");
	links := GetLinksByUser(id)
	flag := false
	for _, link := range links {
		if link.short == "one" && link.long == "two" {
			flag = true
		}
	}
	assert.Equal(t, flag, true)
	deleteUserAndLinks(id)
}

func TestDeleteLink(t *testing.T) {
	info := GetLinksByUser("guest")
	AddLink("guest", "hi", "bye", "unique_id")
	DeleteLink("unique_id")
	info2 := GetLinksByUser("guest")
	assert.Equal(t, len(info), len(info2))
}

