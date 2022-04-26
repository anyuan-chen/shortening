package cockroachdb

import (
	"database/sql"
	"hash/fnv"
	"strconv"

	"github.com/anyuan-chen/urlshortener/server/pkg/shortener"
)

// type LinkRepository interface {
// 	CreateUser(id string) (error)
// 	Get(shortened_link string) (string, error)
// 	Create(shortened_link string, original_link string, user_id string) (Link, error) //for id in table, maybe implement some sort of hash
// 	GetByUserID(user_id string)([]Link, error)
// }

type cockroachLinkRepository struct {
	cockroach *sql.DB
}

func CreateCockroachDB(database_url string) (cockroachLinkRepository, error){
	db, err := sql.Open("postgres", database_url)
	if err != nil {
		return cockroachLinkRepository{}, err
	}
	return cockroachLinkRepository{db}, nil
}

func (c *cockroachLinkRepository) CreateUser(id string) error {
	_, err := c.cockroach.Exec("INSERT INTO users (id) VALUES ($1)", id)
	if err != nil {
		return err
	}
	return nil
}
func (c *cockroachLinkRepository) Get(shortened_link string) (string, error){
	var link shortener.Link
	err := c.cockroach.QueryRow("SELECT * FROM links WHERE shortened_link=$1", shortened_link).Scan(&link.Id, &link.Original_link, &link.Shortened_link, &link.User_id)
	if err != nil {
		return "", err
	}
	return link.Original_link, nil
}
func (c *cockroachLinkRepository) Create(shortened_link string, original_link string, user_id string) (shortener.Link, error){
	h := fnv.New64a()
	h.Write([]byte(shortened_link + original_link + user_id))
	id := strconv.FormatUint(h.Sum64(), 10)
	_, err := c.cockroach.Exec("INSERT INTO links (id, original_link, shortened_link, user_id) VALUES ($1, $2, $3, $4)", id, shortened_link, original_link, user_id)
	if err != nil {
		return shortener.Link{}, err
	}
	return shortener.Link{
		Id: id,
		Original_link: original_link,
		Shortened_link: shortened_link,
		User_id: user_id,
	}, nil
}

func (c *cockroachLinkRepository) GetByUserID(user_id string)([]shortener.Link, error){
	rows, err := c.cockroach.Query("SELECT * FROM links WHERE user_id=$1", user_id)
	var links []shortener.Link
	for rows.Next() {
		var link shortener.Link
		if err := rows.Scan(&link.Id, &link.Original_link, &link.Shortened_link, &link.User_id); err != nil {
			return links, err
		}
		links = append(links, link)
	}
	if err != nil {
		return links, err
	}
	return links, nil
}