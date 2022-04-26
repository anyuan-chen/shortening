package github

import (
	"context"
	"encoding/json"
	"io"

	"github.com/anyuan-chen/urlshortener/server/pkg/shortener"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type OAuthProviderGithub struct {
	config oauth2.Config
	userEndpointURL string
}

func InitializeOAuthProviderGithub(callback_url string, client_id string, client_secret string) OAuthProviderGithub {
	userEndpointURL := "https://api.github.com/user"
	config := &oauth2.Config{
		RedirectURL: callback_url,
		ClientID: client_id,
		ClientSecret: client_secret,
		Endpoint: github.Endpoint,
		Scopes: []string{"user"},
	}
	provider := OAuthProviderGithub{
		config: *config,
		userEndpointURL: userEndpointURL,
	}
	return provider
}

func (o *OAuthProviderGithub) GetLoginRedirect(oauthstate string) string{
	return o.config.AuthCodeURL(oauthstate)
}

func (o *OAuthProviderGithub) CodeExchange(code string)([]byte, error){
	token, err := o.config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	val, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}
	return val, nil
}
func (o *OAuthProviderGithub) GetUserInfo(session shortener.Session)([]byte, error){
	token := oauth2.Token{
		AccessToken: session.Access_token,
		TokenType: session.Token_type,
		Expiry: session.Expiry,
	}
	client := o.config.Client(context.Background(), &token)
	response, err := client.Get(o.userEndpointURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return contents, nil 
}