//Package api provides a HTTP API with authentication and link shortening capabilities.
package api

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
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
//takes queries of the form: http://serverurlhere.end/auth/login?provider=providerhere
func (s *Service ) Login(w http.ResponseWriter, r *http.Request){
	q := r.URL.Query()
	provider := q["provider"]
	if len(provider) != 1{
		http.Error(w, "bad number of url params", http.StatusInternalServerError)
		return
	}
	state := make(map[string]interface{})
	var expiration = time.Now().Add(20 * time.Minute)
	b := make([]byte, 16)
	rand.Read(b)
	state["random"] = b
	state["provider"] = provider[0]
	fmt.Println(provider)
	//b = append(b, []byte(state["provider"].(string))... )\
	jsonByte, err := json.Marshal(state)
	encoded_state := base64.URLEncoding.EncodeToString(jsonByte) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{Name: "oauthstate", Value: encoded_state, Expires: expiration}
	http.SetCookie(w, &cookie)
	redirect_url, err := s.linkService.Login(provider[0], encoded_state)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, redirect_url, http.StatusTemporaryRedirect)
}

//Callback is an HTTP endpoint for the OAuth providers once they have logged
//into the platform. Not meant to be used outside of the callbacks from OAuth providers.
func (s *Service ) Callback(w http.ResponseWriter, r *http.Request){
	oauthstate, _ := r.Cookie("oauthstate")
	if r.FormValue("state") != oauthstate.Value{
		http.Error(w, "bad oauth state", http.StatusInternalServerError)
		return
	}
	var stateData map[string]interface{}
	decoded_base64, err := base64.StdEncoding.DecodeString(r.FormValue("state"))
	//fmt.Println(decoded_base64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(decoded_base64, &stateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//fmt.Println(string(decoded_base64))
	provider := stateData["provider"].(string)
	code := r.FormValue("code")
	fmt.Println("code: " + code, "provider: " + provider )
	token, err := s.linkService.Callback(provider, code)
	fmt.Println("made it after callback")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//fmt.Println(token)
	session_id, err := s.linkService.CreateSession(token.AccessToken, token.RefreshToken, token.TokenType, token.Expiry, stateData["provider"].(string))
	fmt.Println("made it after session_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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