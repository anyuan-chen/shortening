package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig = &oauth2.Config{
	RedirectURL: "http://localhost:8080/auth/google/callback",
	ClientID: os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint: google.Endpoint,
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo"

func GenerateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)
	return state
}

func GetUserComingFromGoogle(code string) ([]byte, error) {
	token, _ := GoogleOauthConfig.Exchange(context.Background(), code)
	// if err != nil {
	// 	return nil, fmt.Errorf("code exchange gone wrong!! %s", err.Error())
	// }
	// response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to retrieve user info %s", err.Error())
	// }
	// defer response.Body.Close()
	// contents, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to read response %s", err.Error())
	// }
	val, _ := json.Marshal(token);
	fmt.Println(val)
	return val, nil

}

func GetGoogleUserInfo (client *http.Client) ([]byte, error){
	response, err := client.Get(oauthGoogleUrlAPI)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user info %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response %s", err.Error())
	}
	return contents, nil
}

func GetGoogleToken (session *sessions.Session) (oauth2.Token, error) {
	expiryTime, err := time.Parse(time.RFC3339, session.Values["expiry"].(string))
	if err != nil {
		log.Fatal(err)
	}
	token := &oauth2.Token{
		AccessToken: session.Values["access_token"].(string),
		TokenType: session.Values["token_type"].(string),
		RefreshToken: session.Values["refresh_token"].(string),
		Expiry: expiryTime,
	}
	return *token, nil
}