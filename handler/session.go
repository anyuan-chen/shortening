package handler

import (
	"os"
	"github.com/gorilla/sessions"
)

// type oauth_session struct {
// 	provider string
// 	access_token string
// 	refresh_token string
// 	expiry string
// }

var sessionStore = sessions.NewFilesystemStore("", []byte(os.Getenv("SESSION_KEY")))