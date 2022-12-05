package deployments

import (
	"fmt"
	"os"
	"time"

	"github.com/pterm/pterm"
	"github.com/wandb/server-cli/pkg/api/deploy"
	"github.com/wandb/server-cli/pkg/api/wandb"
	"github.com/xeonx/timeago"
)

func Licensing() {
	pterm.DefaultSection.Println("Licenses")

	i := GetInstance()
	if i.GetType() == ManagedDedicatedCloud {
		pterm.Success.Print("No licensed required.")
		pterm.Println()
		return
	}

	deploymentID := i.GetDeploymentID()
	if deploymentID != "" {
		d, err := deploy.GetDeployment(deploymentID)
		pterm.Fatal.PrintOnError(err)

		pterm.Success.Print("Licensed instance found: ")
		pterm.Bold.Println(pterm.Green(d.Name))
		return
	}

	licenses, err := deploy.GetActiveLicenseOrders()
	pterm.Fatal.PrintOnError(err)

	if len(licenses) == 0 {
		pterm.DefaultParagraph.Println(
			"We could not find a license assoicated with your account." +
				" Please make sure you are using your work email.",
		)
		pterm.Fatal.Println("No licenses found.")
		os.Exit(1)
	}

	pterm.Println("We found a pending license!")
	pterm.Println(pterm.Gray("If this information is not correct please contact sales."))
	license := licenses[0]

	layout := "2006-01-02T15:04:05.000Z"
	expiresAt, _ := time.Parse(layout, license.ExpiresAt)
	td := [][]string{
		{"Max Users", fmt.Sprint(license.MaxUsers)},
		{"Max Teams", fmt.Sprint(license.MaxTeams)},
		{"Expiration", timeago.English.Format(expiresAt)},
	}
	pterm.DefaultTable.WithBoxed().
		WithStyle(&pterm.ThemeDefault.DebugMessageStyle).
		WithData(td).
		Render()

	pterm.Println()

	viewer, _ := wandb.GetViewer()
	organizations := viewer.Organizations
	createOrganization := pterm.Green("+ Create Organization")

	var organizationID string
	if len(organizations) != 0 {
		organizationNames := []string{createOrganization}
		for _, o := range organizations {
			organizationNames = append(organizationNames, o.Name)
		}

		organizationName, _ := pterm.DefaultInteractiveSelect.
			WithDefaultText("Select organization").
			WithOptions(organizationNames).
			Show()

		for _, o := range organizations {
			if o.Name == organizationName {
				organizationID = o.ID
			}
		}
	}

	if organizationID == "" {
		orgName, _ := pterm.DefaultInteractiveTextInput.
			Show("Organization name")
		org, _ := wandb.CreateOrganization(orgName)
		organizationID = org.ID
	}

	pterm.DefaultParagraph.Println(
		"If you have multiple instances names are useful " +
			"for identify each one.",
	)
	name, err := pterm.DefaultInteractiveTextInput.Show("Instance name")
	pterm.Fatal.PrintOnError(err)
	d, err := deploy.CreateDeploymentFromOrder(license.ID, name, organizationID)
	pterm.Fatal.PrintOnError(err)
	GetInstance().SetDeploymentID(d.ID).Write()
	pterm.Println()
}
