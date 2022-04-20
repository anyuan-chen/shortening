package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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


func GetUserComingFromGithub(code string) (string, error) {
    // Set us the request body as JSON
    requestBodyMap := map[string]string{
        "client_id": GithubOAuthConfig.ClientID,
        "client_secret": GithubOAuthConfig.ClientSecret,
        "code": code,
    }
    requestJSON, _ := json.Marshal(requestBodyMap)

    // POST request to set URL
    req, reqerr := http.NewRequest(
        "POST",
        "https://github.com/login/oauth/access_token",
        bytes.NewBuffer(requestJSON),
    )
    if reqerr != nil {
        return "", errors.New(reqerr.Error())
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")

    // Get the response
    resp, resperr := http.DefaultClient.Do(req)
    if resperr != nil {
        return "", errors.New(resperr.Error())
    }

    // Response body converted to stringified JSON
    respbody, _ := ioutil.ReadAll(resp.Body)

    // Represents the response received from Github
    type githubAccessTokenResponse struct {
        AccessToken string `json:"access_token"`
        TokenType   string `json:"token_type"`
        Scope       string `json:"scope"`
    }

    // Convert stringified JSON to a struct object of type githubAccessTokenResponse
    var ghresp githubAccessTokenResponse
    json.Unmarshal(respbody, &ghresp)

    // Return the access token (as the rest of the
    // details are relatively unnecessary for us)
    return ghresp.AccessToken, nil
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