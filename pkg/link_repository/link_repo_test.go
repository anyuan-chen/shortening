package linkrepository_test

import (
	"hash/fnv"
	"os"
	"strconv"
	"testing"

	"github.com/anyuan-chen/urlshortener/server/pkg/link_repository/cockroachdb"
	useridsha256 "github.com/anyuan-chen/urlshortener/server/pkg/short_link_creator/user_id_sha256"
	"github.com/stretchr/testify/assert"
)

//specify type of package that needs to be tested here, can turn this into a table if multiple need to be tested
var test cockroachdb.CockroachLinkRepository

func init(){
	var err error
	test, err = cockroachdb.CreateCockroachDB(os.Getenv("COCKROACH_DB_DATABASE_URL"))
	if err != nil {
		panic("failed to even initialize")
	}
}

func TestAddUser(t *testing.T) {
	id := "my_test_id"
	test.CreateUser(id)
	returned_id, err := test.GetUser(id)
	assert.Equal(t, id, returned_id)
	assert.Nil(t, err)
	test.DeleteUser(id)
}

func TestCreateAndGetLink(t *testing.T) {
	original_link := "https://google.com"
	user_id := "guest"
	link_shortener := useridsha256.ShortLinkCreator{}
	shortened_link := link_shortener.GenerateShortLink(original_link, user_id)
	test.Create(shortened_link, original_link, user_id)
	original_from_db, err := test.Get(shortened_link)
	assert.Equal(t, original_from_db, original_link)
	assert.Nil(t, err)
	err = test.DeleteLink(GetId(shortened_link, original_link, user_id))
	assert.Nil(t, err)
}

func TestGetByUserID(t *testing.T) {
	test.CreateUser("testgetbyuserid")
	test.Create("short", "original", "testgetbyuserid")
	test.Create("short2", "original2", "testgetbyuserid")
	test.Create("short3", "original3", "testgetbyuserid")
	shorts := map[string]string {"short" : "original", "short2" : "original2", "short3" : "original3"}
	links, err := test.GetByUserID("testgetbyuserid")
	assert.Nil(t, err)
	assert.Equal(t, len(links), 3)
	for _, link := range links {
		assert.Equal(t, link.Original_link, shorts[link.Shortened_link])
	}
	test.DeleteLink(GetId("short", "original", "testgetbyuserid"))
	test.DeleteLink(GetId("short2", "original2", "testgetbyuserid"))
	test.DeleteLink(GetId("short3", "original3", "testgetbyuserid"))
	test.DeleteUser("testgetbyuserid")
}

func GetId(shortened_link string, original_link string, user_id string) string{
	h := fnv.New64a()
	h.Write([]byte(shortened_link + original_link + user_id))
	id := strconv.FormatUint(h.Sum64(), 10)
	return id
}