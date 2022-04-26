package shortener

type OAuthProvider interface {
	GetLoginRedirect() (string) 
	CodeExchange(code string) ([]byte, error)
	GetUserInfo(session Session) ([]byte, error)
}