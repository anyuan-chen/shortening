package shortener

type OAuthProvider interface {
	GetLoginRedirect(oauthstate string) (string) 
	CodeExchange(code string) ([]byte, error)
	GetUserInfo(session Session) ([]byte, error)
}