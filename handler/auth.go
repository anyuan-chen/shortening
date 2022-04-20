package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anyuan-chen/urlshortener/server/auth"
)

func OauthGoogleCallback(w http.ResponseWriter, r *http.Request){
	oauthState, _ := r.Cookie("oauthstate")
	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth state - this has been tampered with")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	data , err := auth.GetUserComingFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error()) 
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "UserInfo: %s\n", data)
}

func OauthGithubCallback(w http.ResponseWriter, r *http.Request) {
	oauthState, _ := r.Cookie("oauthstate")
	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth state - this has been tampered with")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	access_token , err := auth.GetUserComingFromGithub(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error()) 
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	data, err := auth.GetGithubData(access_token)
	if err != nil {
		log.Println(err.Error()) 
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "UserInfo: %s\n", data)
}

func OauthGoogleLogin(w http.ResponseWriter, r *http.Request){
	oauthState := auth.GenerateStateOauthCookie(w)
	u := auth.GoogleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func OauthGithubLogin(w http.ResponseWriter, r *http.Request) {
	oauthState := auth.GenerateStateOauthCookie(w)
	u := auth.GithubOAuthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}