//Package inmemory implements a in-memory store for sessions.
package inmemory

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/anyuan-chen/urlshortener/server/pkg/oauth_provider/github"
	"github.com/anyuan-chen/urlshortener/server/pkg/oauth_provider/google"
	"github.com/anyuan-chen/urlshortener/server/pkg/shortener"
	"github.com/google/uuid"
)

type memorySessionRepository struct {
	sessionStore map[string] shortener.Session	
}
//GetSession takes a session id and returns any active session with that id
func (s *memorySessionRepository) GetSession(session_id string) (shortener.Session, error){
	if s.sessionStore[session_id] == (shortener.Session{}){
		return shortener.Session{}, errors.New("session not found")
	}
	return s.sessionStore[session_id], nil
}	

//GetId takes a session id and attempts to make a call to the OAuth provider to return a user id associated with 
//the session.
func (s *memorySessionRepository) GetId(session_id string) (string, error){
	session, err := s.GetSession(session_id)
	if err != nil {
		return "", err
	}
	callback_url := os.Getenv("REDIRECT_URL")
	client_secret := os.Getenv("OAUTH_CLIENT_SECRET" + strings.ToUpper(session.Provider))
	client_id := os.Getenv("OAUTH_CLIENT_ID" + strings.ToUpper(session.Provider))
	var provider shortener.OAuthProvider
	//bad hardcoding
	if session.Provider == "google"{
		temp := google.InitializeOAuthProvider(callback_url, client_id, client_secret)
		provider = &temp
	} else {
		temp := github.InitializeOAuthProvider(callback_url, client_id, client_secret)
		provider = &temp
	}
	data, err := provider.GetUserInfo(session)
	if err != nil {
		return "", err
	}
	var user_info map[string]interface{}
	err = json.Unmarshal(data, &user_info)
	if err != nil {
		return "", err
	}
	return user_info["id"].(string), nil
}

//IsLoggedIn takes a session id and makes a call to the OAuth provider to return if the session is valid
func (s *memorySessionRepository) IsLoggedIn(session_id string) (bool, error){
	_ , err := s.GetId(session_id)
	if err != nil {
		return false, err
	}
	return true, nil
}

//CreateSession takes in all the parameters for a session, creates it, then stores it in the sessionStore
func (s *memorySessionRepository) CreateSession(access_token string, refresh_token string, token_type string, expiry time.Time, provider string)(string, error){
	session := shortener.Session{Access_token: access_token, Refresh_token: refresh_token, Token_type: token_type, Expiry: expiry, Provider: provider}
	var session_id string
	for condition := true; condition; {
		session_id = uuid.New().String()
		if emptySession := (shortener.Session{}); s.sessionStore[session_id] != emptySession  {
			condition = false
			break;
		}
	}
	s.sessionStore[session_id] = session
	return session_id, nil
}