package wandb

import (
	"encoding/base64"
	"net/http"

	"github.com/hasura/go-graphql-client"
	"github.com/spf13/viper"
)

func AddAPIKey(r *http.Request) {
	apikey := viper.GetString("wandb.apikey")
	auth := "apikey:" + apikey
	hash := base64.StdEncoding.EncodeToString([]byte(auth))
	r.Header.Add("Authorization", "Basic "+hash)
}

var client = graphql.
	NewClient("https://api.wandb.ai/graphql", nil).
	WithRequestModifier(AddAPIKey)
