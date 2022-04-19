package main

import (
	"net/http"

	"github.com/anyuan-chen/urlshortener/server/handler"
	"github.com/anyuan-chen/urlshortener/server/store"
	"github.com/gorilla/mux"
)

func main() {
	
	r := mux.NewRouter()
	store.InitializeStore()
	r.HandleFunc("/create/{url}", handler.CreateShortUrl).Methods("POST")
	r.HandleFunc("/redirect/{url}", handler.RedirectURL).Methods("GET")
	r.HandleFunc("/auth/google/login", handler.OauthGoogleLogin)
	r.HandleFunc("/auth/google/callback", handler.OauthGoogleCallback)
	http.ListenAndServe(":8080", r)
}