package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/anyuan-chen/urlshortener/server/auth"
	"github.com/anyuan-chen/urlshortener/server/shortener"
	"github.com/anyuan-chen/urlshortener/server/store"
)

//this handles the /create



func CreateShortUrl(w http.ResponseWriter, r *http.Request){
	// /create/shortlinkhere
	type response struct {
		Url string `json:"url"`
	}
	pathElements := strings.Split(r.URL.Path, "/")
	params := r.URL.Query()
	if len(params["user_id"]) == 0 {
		http.Error(w, "invalid request - you may not proceed without a user id", http.StatusInternalServerError)
	}
	shortenedUrl := shortener.GenerateShortLink(pathElements[len(pathElements) - 1], params["user_id"][0])
	store.InsertUrl(shortenedUrl, pathElements[len(pathElements)-1], params["user_id"][0])

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response{shortenedUrl})
}

func Redirect(w http.ResponseWriter, r *http.Request){
	type response struct {
		Url string `json:"url"`
		Err error `json:"error"`
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	pathElements := strings.Split(r.URL.Path, "/")
	shortenedUrl, err := store.RetrieveUrl(pathElements[len(pathElements) -1])
	if err != nil {
		json.NewEncoder(w).Encode(response{Url: "", Err: err})
	} else {
		json.NewEncoder(w).Encode(response{Url: shortenedUrl, Err: nil})
	}
}

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

	fmt.Fprintf(w, "%s", data)
}

func OauthGoogleLogin(w http.ResponseWriter, r *http.Request){
	oauthState := auth.GenerateStateOauthCookie(w)
	u := auth.GoogleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}