package token

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

var TokenConfig *clientcredentials.Config
var Client *http.Client

func SetUp(clientID string, clientSecret string) {
	TokenConfig = &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://osu.ppy.sh/oauth/token",
		Scopes:       []string{"public"},
	}
	Client = TokenConfig.Client(context.Background())
}
