package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Redirect takes a URL which was the result of a previous shortening operation,
//then redirects the user to the original URL.
func (s *Service ) Redirect(w http.ResponseWriter, r *http.Request){
	shortened_link := mux.Vars(r)["url"]
	original_link, err := s.linkService.Get(shortened_link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, original_link, http.StatusPermanentRedirect)
}
//CreateAuthenticated is meant as a way for logged in users to shorten a link.
//This will then be accessible to them if they use the GetLinksForUserID endpoint
func (s *Service ) CreateAuthenticated(w http.ResponseWriter, r *http.Request){
	
}
//CreateUnauthenticated is a meant as a way for unauthenticated users to shorten a link.
func (s *Service ) CreateUnauthenticated(){

}
//GetLinksForUserID returns all links created by a specific user from the CreateAuthenticated
//handler.
func (s *Service ) GetLinksForUserID(){

}