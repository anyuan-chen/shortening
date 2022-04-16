package main

import (
	"github.com/anyuan-chen/urlshortener/server/handler"
	"github.com/anyuan-chen/urlshortener/server/store"
	"github.com/anyuan-chen/urlshortener/server/util"
	"github.com/gorilla/mux"
)

func main() {
	util.LoadEnv()
	r := mux.NewRouter()
	store.InitializeStore()
	r.HandleFunc("/create/{url}", handler.CreateShortUrl)
	r.HandleFunc("/redirect/{url}", handler.Redirect)
}