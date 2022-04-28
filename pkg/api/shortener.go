package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anyuan-chen/urlshortener/server/pkg/shortener"
	"github.com/gorilla/mux"
)

//Redirect takes a URL which was the result of a previous shortening operation,
//then redirects the user to the original URL.
func (s *Service ) Redirect(w http.ResponseWriter, r *http.Request){
	shortened_link := mux.Vars(r)["url"]
	fmt.Println(shortened_link, "original_link")
	original_link, err := s.linkService.Get(shortened_link)
	fmt.Println(err, "redirect_link_service")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(original_link))
}
//Create is meant as a way for logged in users to shorten a link. If a user id id provided,
//it will be accessible to them if they use the GetLinksForUserID endpoint. If not, 
//they will have to save the link on their own for future usage
func (s *Service ) Create(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	id := r.Context().Value("id")
	original_url_params := r.URL.Query()["original_url"]
	var link shortener.Link
	var err error
	if len(original_url_params) != 1 {
		http.Error(w, "bad query parameters", http.StatusBadRequest)
		return;
	}  
	if id != nil{
		link, err = s.linkService.CreateAuthenticated(original_url_params[0], id.(string))
	} else {
		link, err = s.linkService.CreateUnauthenticated(original_url_params[0])
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}
	link_json, err := json.Marshal(link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(link_json)
}

//GetLinksForUserID returns all links created by a specific user from the CreateAuthenticated
//handler.
func (s *Service ) GetLinksForUserID(w http.ResponseWriter, r *http.Request){
	session_id := r.Context().Value("session_id")
	links, err := s.linkService.GetByUserID(session_id.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}
	links_json, err := json.Marshal(links)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}
	defer r.Body.Close()
	r.Body.Read(links_json)
}