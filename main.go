package main

import (
	"net/http"

	"github.com/anyuan-chen/urlshortener/server/handler"
	"github.com/anyuan-chen/urlshortener/server/store"
	"github.com/anyuan-chen/urlshortener/server/users"
	"github.com/gorilla/mux"
)

func main() {
	
	r := mux.NewRouter()
	store.InitializeStore()
	users.InitializeDatabase()
	
	r.HandleFunc("/create/{url}", handler.CreateShortUrl).Methods("POST")
	r.HandleFunc("/redirect/{url}", handler.RedirectURL).Methods("GET")
	r.HandleFunc("/auth/google/login", handler.OauthGoogleLogin)
	r.HandleFunc("/auth/google/callback", handler.OauthGoogleCallback)
	r.HandleFunc("/auth/github/login", handler.OauthGithubLogin)
	r.HandleFunc("/auth/github/callback", handler.OauthGithubCallback)
	r.Handle("/id", handler.LoggedInMiddleware(handler.GetUser))
	r.Handle("/links", handler.LoggedInMiddleware(handler.GetLinksForUser))
	http.ListenAndServe(":8080", r)
}