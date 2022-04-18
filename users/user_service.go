package users

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/cockroachdb/cockroach-go/crdb"
)

// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

type Database struct {
	db *sql.DB
}

var user_service = &Database{}

func InitializeDatabase() *Database {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	user_service.db = db;
	return user_service
}

func GetUser (id string) (string, string) {
	rows, err := user_service.db.Query("SELECT * FROM users WHERE id= $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var googleid, profileUrl string
	var count int
	for rows.Next(){
		if err := rows.Scan(&googleid, &profileUrl); err != nil {
			log.Fatal(err)
		}
		count++

	}
	if count == 0 {
		log.Fatal("user doesn't exist")
	} else if count > 1 {
		log.Fatal("more than user with this id")
	}
	return id, profileUrl
}

func CreateUser ( id string, profileUrl string ) error {
	executeQuery := func (tx *sql.Tx, id string, pfpurl string) error {
		if _, err := tx.Exec ("INSERT INTO users (name, 'profile url') VALUES ($1, $2)", id, profileUrl); err != nil {
			return err
		}
		return nil
	}
	err := crdb.ExecuteTx(context.Background(), user_service.db, nil, func(tx *sql.Tx) error { 
		return executeQuery(tx, id, profileUrl);
	})
	if err == nil {
		return nil
	}
	return err
} 

func GetLinksByUser (id string) map[string]string {
	rows, err := user_service.db.Query("SELECT * FROM links WHERE userid=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	linkMap := make(map[string]string)
	for rows.Next() {
		var short, long string
		if err := rows.Scan(&short, &long); err != nil {
			log.Fatal(err)
		}
		linkMap[short] = long
	}
	return linkMap
}