package users

import (
	"context"
	"database/sql"
	"log"

	"github.com/cockroachdb/cockroach-go/crdb"
)

func GetUser (db *sql.DB, id string) (string, string) {
	rows, err := db.Query("SELECT * FROM users WHERE id= $1", id)
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

func CreateUser (db *sql.DB, id string, profileUrl string ) error {
	executeQuery := func (tx *sql.Tx, id string, pfpurl string) error {
		if _, err := tx.Exec ("INSERT INTO users (name, 'profile url') VALUES ($1, $2)", id, profileUrl); err != nil {
			return err
		}
		return nil
	}
	err := crdb.ExecuteTx(context.Background(), db, nil, func(tx *sql.Tx) error { 
		return executeQuery(tx, id, profileUrl);
	})
	if err == nil {
		return nil
	}
	return err
} 

func GetLinksByUser (db *sql.DB, id string) map[string]string {
	rows, err := db.Query("SELECT * FROM links WHERE userid=$1", id)
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