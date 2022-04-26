package google

import (
	"context"
	"encoding/json"
	"io"

	"github.com/anyuan-chen/urlshortener/server/pkg/shortener"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


type OAuthProviderGoogle struct {
	config oauth2.Config
	userEndpointURL string
}

func InitializeOAuthProvider(callback_url string, client_id string, client_secret string) OAuthProviderGoogle {
	userEndpointURL := "https://www.googleapis.com/oauth2/v2/userinfo"
	config := &oauth2.Config{
		RedirectURL: callback_url,
		ClientID: client_id,
		ClientSecret: client_secret,
		Endpoint: google.Endpoint,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
	provider := OAuthProviderGoogle{
		config: *config,
		userEndpointURL: userEndpointURL,
	}
	return provider
}

func (o *OAuthProviderGoogle) GetLoginRedirect(oauthstate string) string{
	return o.config.AuthCodeURL(oauthstate, oauth2.AccessTypeOffline)
}

func (o *OAuthProviderGoogle) CodeExchange(code string)([]byte, error){
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

func (o *OAuthProviderGoogle) GetUserInfo(session shortener.Session)([]byte, error){
	token := oauth2.Token{
		AccessToken: session.Access_token,
		RefreshToken: session.Refresh_token,
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