package cockroachdb

import (
	"database/sql"
	"hash/fnv"
	"strconv"

	"github.com/anyuan-chen/urlshortener/server/pkg/shortener"
)

type CockroachLinkRepository struct {
	cockroach *sql.DB
}

func CreateCockroachDB(database_url string) (CockroachLinkRepository, error){
	db, err := sql.Open("postgres", database_url)
	if err != nil {
		return CockroachLinkRepository{}, err
	}
	return CockroachLinkRepository{db}, nil
}

func (c *CockroachLinkRepository) CreateUser(id string) error {
	_, err := c.cockroach.Exec("INSERT INTO users (id) VALUES ($1)", id)
	if err != nil {
		return err
	}
	return nil
}
func (c *CockroachLinkRepository) Get(shortened_link string) (string, error){
	var link shortener.Link
	err := c.cockroach.QueryRow("SELECT * FROM links WHERE shortened_link=$1", shortened_link).Scan(&link.Id, &link.Original_link, &link.Shortened_link, &link.User_id)
	if err != nil {
		return "", err
	}
	return link.Original_link, nil
}
func (c *CockroachLinkRepository) Create(shortened_link string, original_link string, user_id string) (shortener.Link, error){
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

func (c *CockroachLinkRepository) GetByUserID(user_id string)([]shortener.Link, error){
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