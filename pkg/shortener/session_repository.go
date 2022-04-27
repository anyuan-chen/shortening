package shortener

import (
	"time"

	"golang.org/x/oauth2"
)

type SessionRepository interface {
	CreateSessionRepository()
	GetSession(session_id string) (Session, error)
	GetId(session_id string) (string, error)
	IsLoggedIn(session_id string) (bool, error)
	CreateSession(access_token string, refresh_token string, token_type string, expiry time.Time, provider string) (string, error) 
	GetLoginRedirect(provider string, oauthstate string) (string, error)
	CodeExchange(provider string, code string) (*oauth2.Token, error)
}