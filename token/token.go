package token

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

var __token_config *clientcredentials.Config
var Client *http.Client

func SetUp(clientID string, clientSecret string) {
	__token_config = &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://osu.ppy.sh/oauth/token",
		Scopes:       []string{"public"},
	}
	Client = __token_config.Client(context.Background())
}
