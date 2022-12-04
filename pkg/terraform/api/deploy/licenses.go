package deploy

import (
	"context"

	"github.com/hasura/go-graphql-client"
	"github.com/wandb/server-cli/pkg/terraform/api/wandb"
)

type License struct {
	ID      string
	License string
	Trial   bool
}

func GetLicense(deploymentID string) (string, error) {
	deployment, err := GetDeployment(deploymentID)
	if len(deployment.Licenses.Nodes) > 0 {
		return deployment.Licenses.Nodes[0].License, err
	}
	return "", err
}

func ClaimLicenseOrder() (string, error) {
	return "", nil
}

type LicenseOrders struct {
	ID    string
	Flags []struct {
		Flag        string
		Name        string
		Description string
	}
	MaxStorageGb  int
	MaxUsers      int
	MaxTeams      int
	CustomerEmail string
	ExpiresAt     string
	CreatedAt     string
}

func GetActiveLicenseOrders() ([]LicenseOrders, error) {
	viewer, err := wandb.GetViewer()
	if err != nil {
		return nil, err
	}
	var query struct {
		LicenseOrders struct {
			Nodes []LicenseOrders
		} `graphql:"licenseOrders(first: $first, where: { customerEmail: { equals: $email }, deploymentId: { equals: null } })"`
	}
	variables := map[string]interface{}{
		"first": graphql.NewInt(50),
		"email": viewer.Email,
	}
	err = client.Query(
		context.Background(),
		&query,
		variables,
		graphql.OperationName("ServerCLIOrders"),
	)
	return query.LicenseOrders.Nodes, err
}
