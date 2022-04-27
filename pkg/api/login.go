//Package api provides a HTTP API with authentication and link shortening capabilities.
package api

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/anyuan-chen/urlshortener/server/pkg/shortener"
)

type Service struct {
	linkService shortener.LinkService
}

func NewService(linkService shortener.LinkService) Service{
	return Service{linkService: linkService}
}

//Login is meant as an HTTP endpoint for users to login into the platform.
//This endpoint redirects to the user-parameter specified OAuth endpoint.
func (s *Service ) Login(w http.ResponseWriter, r *http.Request){
	url, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, "error reading your url", http.StatusInternalServerError)
	}
	q := url.Query()
	provider := q["provider"]
	if len(provider) != 1{
		http.Error(w, "bad number of url params", http.StatusInternalServerError)
	}
	state := make(map[string]interface{})
	var expiration = time.Now().Add(20 * time.Minute)
	b := make([]byte, 16)
	rand.Read(b)
	state["random"] = base64.URLEncoding.EncodeToString(b) 
	state["provider"] = provider
	jsonByte, err := json.Marshal(state)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	state_string := string(jsonByte)
	cookie := http.Cookie{Name: "oauthstate", Value: state_string, Expires: expiration}
	http.SetCookie(w, &cookie)
	redirect_url, err := s.linkService.Login(provider[0], state_string)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, redirect_url, http.StatusTemporaryRedirect)
}

//Callback is an HTTP endpoint for the OAuth providers once they have logged
//into the platform. 
func (s *Service ) Callback(w http.ResponseWriter, r *http.Request){
	oauthstate, _ := r.Cookie("oauthstate")
	if r.FormValue("state") != oauthstate.Value{
		http.Error(w, "bad oauth state", http.StatusInternalServerError)
	}
	var stateData map[string]interface{}
	err := json.Unmarshal([]byte(r.FormValue("state")), &stateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	token, err := s.linkService.Callback(stateData["provider"].(string), r.FormValue("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	session_id, err := s.linkService.CreateSession(token.AccessToken, token.RefreshToken, token.TokenType, token.Expiry, stateData["provider"].(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	session_cookie := http.Cookie{
		Name: "session_id",
		Path: "/",
		Value: session_id,
		HttpOnly: false,
		Expires: time.Now().Add(time.Hour),
	}
	http.SetCookie(w, &session_cookie)
	http.Redirect(w, r, os.Getenv("FRONTEND_URL") + "/dashboard", http.StatusTemporaryRedirect)
}