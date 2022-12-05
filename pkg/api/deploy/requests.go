package deploy

import (
	"github.com/hasura/go-graphql-client"
	"github.com/wandb/server-cli/pkg/api/wandb"
)

var client = graphql.
	NewClient("https://deploy.wandb.ai/api/graphql", nil).
	WithRequestModifier(wandb.AddAPIKey)
