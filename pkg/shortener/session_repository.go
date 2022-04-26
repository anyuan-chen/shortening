package shortener

type SessionRepository interface {
	CreateSessionRepository()
	GetSession(session_id string) (Session, error)
	GetId(session_id string) (string, error)
	IsLoggedIn(session_id string) (Link, error)
	CreateSession(access_token string, refresh_token string, token_type string, expiry string, provider string) (string, error) 
	GetLoginRedirect(provider string, oauthstate string) (string, error)
	CodeExchange(provider string, code string) ([]byte, error)
}