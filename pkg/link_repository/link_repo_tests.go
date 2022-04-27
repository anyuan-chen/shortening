package linkrepository_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/anyuan-chen/urlshortener/server/pkg/link_repository/cockroachdb"
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
	assert.Nil(t, err)
	assert.Equal(t, id, returned_id)
}