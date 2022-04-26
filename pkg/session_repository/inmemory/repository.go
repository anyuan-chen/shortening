//Package inmemory implements a in-memory store for sessions.
package inmemory

import (
	"errors"
	"time"

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
	return "", errors.New("not imp")
}

//IsLoggedIn takes a session id and makes a call to the OAuth provider to return if the session is valid
func (s *memorySessionRepository) IsLoggedIn(session_id string) (bool, error){
	return false, errors.New("not imp")
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