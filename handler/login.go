package handler

import (
	"net/http"

	"github.com/anyuan-chen/urlshortener/server/auth"
	"golang.org/x/oauth2"
)
func OauthGoogleLogin(w http.ResponseWriter, r *http.Request){
	oauthState := auth.GenerateStateOauthCookie(w)
	u := auth.GoogleOauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func OauthGithubLogin(w http.ResponseWriter, r *http.Request) {
	oauthState := auth.GenerateStateOauthCookie(w)
	u := auth.GithubOAuthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}