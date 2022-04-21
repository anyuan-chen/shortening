package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var GithubOAuthConfig = &oauth2.Config{
	RedirectURL: "http://localhost:8080/auth/github/callback",
	ClientID: os.Getenv("THE_GITHUB_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("THE_GITHUB_OAUTH_CLIENT_SECRET"),	
	Endpoint: github.Endpoint,
	Scopes: []string{"user"},
}



func GetUserComingFromGithub(code string) ([]byte, error) {
    token, err := GithubOAuthConfig.Exchange(context.Background(), code)
    if err != nil {
        return nil, err
    }
    val, err := json.Marshal(token)
    if err != nil {
        return nil, err
    }
    return val, nil

    // Set us the request body as JSON
    // requestBodyMap := map[string]string{
    //     "client_id": GithubOAuthConfig.ClientID,
    //     "client_secret": GithubOAuthConfig.ClientSecret,
    //     "code": code,
    // }
    // requestJSON, _ := json.Marshal(requestBodyMap)

    // // POST request to set URL
    // req, reqerr := http.NewRequest(
    //     "POST",
    //     "https://github.com/login/oauth/access_token",
    //     bytes.NewBuffer(requestJSON),
    // )
    // if reqerr != nil {
    //     return "", errors.New(reqerr.Error())
    // }
    // req.Header.Set("Content-Type", "application/json")
    // req.Header.Set("Accept", "application/json")

    // // Get the response
    // resp, resperr := http.DefaultClient.Do(req)
    // if resperr != nil {
    //     return "", errors.New(resperr.Error())
    // }

    // // Response body converted to stringified JSON
    // respbody, _ := ioutil.ReadAll(resp.Body)

    // // Represents the response received from Github
    // type githubAccessTokenResponse struct {
    //     AccessToken string `json:"access_token"`
    //     TokenType   string `json:"token_type"`
    //     Scope       string `json:"scope"`
    // }

    // // Convert stringified JSON to a struct object of type githubAccessTokenResponse
    // var ghresp githubAccessTokenResponse
    // json.Unmarshal(respbody, &ghresp)

    // // Return the access token (as the rest of the
    // // details are relatively unnecessary for us)
    // return ghresp.AccessToken, nil
}
const oauthGithubUrlAPI = "https://api.github.com/user"
func GetGithubUserInfo(client *http.Client) ([]byte, error) {
    // Get request to a set URL
    response, err := client.Get(oauthGithubUrlAPI)
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

func GetGithubData(accessToken string) (string, error) {
    // Get request to a set URL
    req, reqerr := http.NewRequest(
        "GET",
        "https://api.github.com/user",
        nil,
    )
    if reqerr != nil {
        return "", errors.New(reqerr.Error())
    }

    // Set the Authorization header before sending the request
    // Authorization: token XXXXXXXXXXXXXXXXXXXXXXXXXXX
    authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
    req.Header.Set("Authorization", authorizationHeaderValue)

    // Make the request
    resp, resperr := http.DefaultClient.Do(req)
    if resperr != nil {
		return "", errors.New(resperr.Error())
    }

    // Read the response as a byte slice
    respbody, _ := ioutil.ReadAll(resp.Body)

    // Convert byte slice to string and return
    return string(respbody), nil
}

func GetGithubToken(session *sessions.Session) (oauth2.Token, error){
    expiryTime, err := time.Parse(time.RFC3339, session.Values["expiry"].(string))
    if err != nil {
        log.Fatal(err)
    }
    token := oauth2.Token{
        AccessToken: session.Values["access_token"].(string),
        TokenType: session.Values["token_type"].(string),
        Expiry: expiryTime,
    }
    return token, nil
}