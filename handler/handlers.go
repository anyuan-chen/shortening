package handler

import (
	"encoding/json"
	"net/http"
	"strings"

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