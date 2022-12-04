package wandb

import (
	"context"

	"github.com/hasura/go-graphql-client"
)

type Organization struct {
	ID   string
	Name string
}

func CreateOrganization(name string) (Organization, error) {
	var mutation struct {
		CreateLocalLicenseOrganization Organization `graphql:"createLocalLicenseOrganiation(input: { newOrganizationName: $name })"`
	}
	variables := map[string]interface{}{
		"name": name,
	}
	err := client.Mutate(
		context.Background(),
		&mutation,
		variables,
		graphql.OperationName("ServerCLICreateOrganization"),
	)
	return mutation.CreateLocalLicenseOrganization, err
}
