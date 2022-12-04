package deploy

import (
	"context"

	"github.com/hasura/go-graphql-client"
)

type Deployment struct {
	ID             string
	Name           string
	OrganizationID string
	Description    string

	Licenses struct {
		Nodes []License
	} `graphql:"licenses(first: 2)"`
}

func GetDeployment(id string) (Deployment, error) {
	var query struct {
		Deployment Deployment `graphql:"deployment(id: $id)"`
	}
	variables := map[string]interface{}{
		"id": id,
	}
	err := client.Query(
		context.Background(),
		&query,
		variables,
		graphql.OperationName("ServerCLIDeployment"),
	)
	return query.Deployment, err
}

type DeploymentOrderCreateInput struct {
	name           string
	description    string
	organizationId string
}

func CreateDeploymentFromOrder(
	licenseOrderID string,
	name string,
	organizationId string,
) (Deployment, error) {
	var mutation struct {
		Deployment Deployment `graphql:"createDeploymentFromOrder(data: { type: MANUAL, name: $name, organizationId: $organizationId, licenseOrderId: $licenseOrderId})"`
	}
	variables := map[string]interface{}{
		"licenseOrderId": licenseOrderID,
		"name":           name,
		"organizationId": graphql.ID(organizationId),
	}
	err := client.Mutate(
		context.Background(),
		&mutation,
		variables,
		graphql.OperationName("ServerCLICreateDeployment"),
	)
	return mutation.Deployment, err
}
