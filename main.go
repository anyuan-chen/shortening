package main

import (
	"net/http"

	"github.com/anyuan-chen/urlshortener/server/auth"
	"github.com/anyuan-chen/urlshortener/server/handler"
	"github.com/anyuan-chen/urlshortener/server/store"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	store.InitializeStore()
	r.HandleFunc("/create/{url}", handler.CreateShortUrl).Methods("POST")
	r.HandleFunc("/redirect/{url}", handler.Redirect).Methods("GET")
	r.HandleFunc("/auth/google/login", auth.OauthGoogleLogin)
	r.HandleFunc("/auth/google/callback", auth.OauthGoogleCallback)
	http.ListenAndServe(":8080", r)

}