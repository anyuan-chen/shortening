package shortener

import (
	"time"

	"golang.org/x/oauth2"
)

type LinkService interface {
	Get(shortened_link string) (string, error)
	CreateAuthenticated(id string, shortened_linkstring , original_link string, user_id string) (Link, error)
	CreateUnauthenticated(id string, shortened_link string, original_link string)(Link, error)
	GetByUserID(session_id string)([]Link, error)
	Login(provider string, oauthstate string) (string, error)
	Callback(provider string, code string) (*oauth2.Token, error)
	CreateSession(access_token string, refresh_token string, token_type string, expiry time.Time, provider string)(string, error)
}

