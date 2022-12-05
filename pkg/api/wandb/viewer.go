package wandb

import (
	"context"
	"errors"

	"github.com/hasura/go-graphql-client"
)

type Viewer struct {
	ID            string
	Name          string
	Email         string
	Admin         bool
	Username      string
	Entity        string
	Organizations []Organization
}

func GetViewer() (*Viewer, error) {
	var query struct {
		Viewer Viewer
	}
	variables := map[string]interface{}{}
	err := client.Query(
		context.Background(),
		&query,
		variables,
		graphql.OperationName("ServerCLIViewer"),
	)
	if query.Viewer.ID == "" {
		err = errors.New("Invalid API key")
	}
	return &query.Viewer, err
}
