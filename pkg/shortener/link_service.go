package shortener

import (
	"time"

	"golang.org/x/oauth2"
)

type LinkService interface {
	Get(shortened_link string) (string, error)
	CreateAuthenticated(original_link string, user_id string) (Link, error)
	CreateUnauthenticated(original_link string)(Link, error)
	GetByUserID(session_id string)([]Link, error)
	Login(provider string, oauthstate string) (string, error)
	Callback(provider string, code string) (*oauth2.Token, error)
	CreateSession(access_token string, refresh_token string, token_type string, expiry time.Time, provider string)(string, error)
	ValidateSession(session_id string) (string, error)
	GetSession( session_id string)(Session, error)
}

