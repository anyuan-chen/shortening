package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/anyuan-chen/urlshortener/server/auth"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
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
	type GoogleOauthResponse struct {
		Access_token string `json:"access_token"`
		Token_type string `json:"token_type"`
		Refresh_token string`json:"refresh_token"`
		Expiry string`json:"expiry"`
	}
	var dataJson GoogleOauthResponse
	err = json.Unmarshal(data, &dataJson)
	if err != nil {
		log.Println(err.Error()) 
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	uuid := uuid.New().String();
	session, err := sessionStore.Get(r, uuid)
	if err != nil {
		log.Println(err.Error()) 
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	session.Values["access_token"] = dataJson.Access_token;
	session.Values["token_type"] = dataJson.Token_type;
	session.Values["refresh_token"] = dataJson.Refresh_token;
	session.Values["expiry"] = dataJson.Expiry;
	session.Values["provider"] = "google"
	err = sessions.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session_cookie := http.Cookie{
		Name: "session_id",
		Path: "/",
		Value: uuid,
		HttpOnly: true,
		Expires: time.Now().Add(time.Hour),
	}
	http.SetCookie(w, &session_cookie)
	http.Redirect(w, r, "/dashboard", http.StatusOK)
}

func OauthGithubCallback(w http.ResponseWriter, r *http.Request) {
	oauthState, _ := r.Cookie("oauthstate")
	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth state - this has been tampered with")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	data , err := auth.GetUserComingFromGithub(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error()) 
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	type GithubOAuthResponse struct {
		Access_token string `json:"access_token"`
		Token_type string `json:"token_type"`
		Expiry string `json:"expiry"`
	}
	var dataJson GithubOAuthResponse
	err = json.Unmarshal(data, &dataJson)
	if err != nil {
		log.Println(err.Error()) 
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	uuid := uuid.New().String()
	session, _ := sessionStore.Get(r, uuid)
	session.Values["access_token"] = dataJson.Access_token
	session.Values["token_type"] = dataJson.Token_type
	session.Values["expiry"] = dataJson.Expiry
	session.Values["provider"] = "github"
	err = sessions.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session_cookie := http.Cookie{
		Name: "session_id",
		Path: "/",
		Value: uuid,
		HttpOnly: true,
		Expires: time.Now().Add(time.Hour),
	}
	http.SetCookie(w, &session_cookie)
	http.Redirect(w, r, "/dashboard", http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {
	session_id, err := r.Cookie("session_id")
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "http://localhost:8080/auth/google/login", http.StatusFound)
		return
	}
	//check if there exists a session
	session, err := sessionStore.Get(r, session_id.Value)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	var client *http.Client
	expiryTime, err := time.Parse(time.RFC3339, session.Values["expiry"].(string))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if session.Values["provider"] == "google" {
		token := &oauth2.Token{
			AccessToken: session.Values["access_token"].(string),
			TokenType: session.Values["token_type"].(string),
			RefreshToken: session.Values["refresh_token"].(string),
			Expiry: expiryTime,
		}
		client = auth.GoogleOauthConfig.Client(context.Background(), token)
		resp, err := auth.GetGoogleUserInfo(client)
		if err != nil {
			log.Println(err.Error())
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		fmt.Print(string(resp))
		if resp != nil {
			http.Redirect(w, r, "https://google.com", http.StatusFound)
		}
	} else if session.Values["provider"] == "github" {
		token := &oauth2.Token{
			AccessToken: session.Values["access_token"].(string),
			TokenType: session.Values["token_type"].(string),
			Expiry: expiryTime,
		}
		client = auth.GithubOAuthConfig.Client(context.Background(), token)
		resp, err := auth.GetGithubUserInfo(client)
		if err != nil {
			log.Println(err.Error())
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		fmt.Print(string(resp))

		if resp != nil {
			http.Redirect(w, r, "https://google.com", http.StatusFound)
		}
	} else { //no session found with that id
		http.Redirect(w, r, "http://localhost:8080/auth/google/login", http.StatusFound)
	}
}

