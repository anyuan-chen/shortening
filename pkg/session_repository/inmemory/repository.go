//Package inmemory implements a in-memory store for sessions.
package inmemory

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/anyuan-chen/urlshortener/server/pkg/oauth_provider/github"
	"github.com/anyuan-chen/urlshortener/server/pkg/oauth_provider/google"
	"github.com/anyuan-chen/urlshortener/server/pkg/shortener"
	"github.com/google/uuid"
)

type MemorySessionRepository struct {
	sessionStore map[string] shortener.Session
	googleOAuth google.OAuthProviderGoogle	
	githubOAuth github.OAuthProviderGithub
}

func (s *MemorySessionRepository) CreateSessionRepository(){
	callback_url := os.Getenv("REDIRECT_URL")
	google_client_secret := os.Getenv("OAUTH_CLIENT_SECRET_GOOGLE")
	google_client_id := os.Getenv("OAUTH_CLIENT_ID_GOOGLE")
	s.googleOAuth = google.InitializeOAuthProvider(callback_url, google_client_id, google_client_secret)
	github_client_secret := os.Getenv("OAUTH_CLIENT_SECRET_GITHUB")
	github_client_id := os.Getenv("OAUTH_CLIENT_ID_GITHUB")
	s.githubOAuth = github.InitializeOAuthProvider(callback_url, github_client_id, github_client_secret)
}

//GetSession takes a session id and returns any active session with that id
func (s *MemorySessionRepository) GetSession(session_id string) (shortener.Session, error){
	if s.sessionStore[session_id] == (shortener.Session{}){
		return shortener.Session{}, errors.New("session not found")
	}
	return s.sessionStore[session_id], nil
}	

//GetId takes a session id and attempts to make a call to the OAuth provider to return a user id associated with 
//the session.
func (s *MemorySessionRepository) GetId(session_id string) (string, error){
	session, err := s.GetSession(session_id)
	if err != nil {
		return "", err
	}
	var data []byte
	if session.Provider == "google"{
		data, err = s.googleOAuth.GetUserInfo(session)
	} else if session.Provider == "github"{
		data, err = s.githubOAuth.GetUserInfo(session)
	} else{
		return "", err
	}
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
func (s *MemorySessionRepository) IsLoggedIn(session_id string) (bool, error){
	_ , err := s.GetId(session_id)
	if err != nil {
		return false, err
	}
	return true, nil
}

//CreateSession takes in all the parameters for a session, creates it, then stores it in the sessionStore
func (s *MemorySessionRepository) CreateSession(access_token string, refresh_token string, token_type string, expiry time.Time, provider string)(string, error){
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

func (s *MemorySessionRepository) GetLoginRedirect(provider string, oauthstate string) (string, error){
	if provider == "google"{
		return s.googleOAuth.GetLoginRedirect(oauthstate), nil
	} else if provider == "github"{
		return s.githubOAuth.GetLoginRedirect(oauthstate), nil
	} else{
		return "", errors.New("invalid provider")
	}
}

func (s *MemorySessionRepository) CodeExchange(provider string, code string) ([]byte, error) {
	var val []byte
	var err error
	if provider == "google"{
		val, err = s.googleOAuth.CodeExchange(code)
		if err != nil {
			return nil, err
		}
	} else if provider == "github"{
		val, err = s.githubOAuth.CodeExchange(code)
		if err != nil {
			return nil, err
		}
	} else{
		return nil, errors.New("invalid provider")
	}
	return val, nil
}