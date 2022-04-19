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
	data, err := auth.GetUserComingFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error()) 
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	http.Redirect(w, r , fmt.Sprintf("http://www.google.com?token=%v", data) + string(data) ,http.StatusFound)
}

func OauthGoogleLogin(w http.ResponseWriter, r *http.Request){
	oauthState := auth.GenerateStateOauthCookie(w)
	u := auth.GoogleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}