package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/anyuan-chen/urlshortener/server/auth"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
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

func LoggedInMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request ) {
		ctx := context.Background()
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
		if session.Values["provider"] == "google" {
			//get the token from the existing session
			token, err := auth.GetGoogleToken(session)
			if err != nil {
				log.Println(err.Error())
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
			//get http client from the session
			client = auth.GoogleOauthConfig.Client(context.Background(), &token)
			//test a request
			resp, err := auth.GetGoogleUserInfo(client)
			if err != nil {
				log.Println(err.Error())
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
			if resp != nil {
				ctx = context.WithValue(ctx, "client", client)
				ctx = context.WithValue(ctx, "provider", "google")
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		} else if session.Values["provider"] == "github" {
			token, _ := auth.GetGithubToken(session)
			client = auth.GithubOAuthConfig.Client(context.Background(), &token)
			resp, err := auth.GetGithubUserInfo(client)
			if err != nil {
				log.Println(err.Error())
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
			if resp != nil {
				ctx = context.WithValue(ctx, "client", client)
				ctx = context.WithValue(ctx, "provider", "github")

				next.ServeHTTP(w, r.WithContext(ctx))
			}
		} 
		 //no session found with that id
		http.Redirect(w, r, "http://localhost:8080/auth/google/login", http.StatusFound)
	})

}

