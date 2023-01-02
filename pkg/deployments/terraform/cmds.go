package terraform

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/pterm/pterm"
)

type TFVersion struct {
	Version string `json:"terraform_version"`
}

var tfVersion TFVersion

func SetTFPath(p string) {

	path := os.Getenv("PATH") + ":" + p
	os.Setenv("PATH", path)

	cmd, err := exec.Command("terraform", "version", "-json").Output()
	if err != nil {
		pterm.Error.Println()
		os.Exit(1)
	}

	json.Unmarshal(cmd, &tfVersion)

	pterm.Println()
	pterm.Info.Printfln("Terraform %s installed", tfVersion.Version)

}

func InstallTerraform(c string) {

	constraint, _ := version.NewConstraint(c)

	tfPath, _ := pterm.DefaultInteractiveTextInput.
		WithDefaultText("Install path (" + os.Getenv("HOME") + "/bin)").
		Show()
	pterm.Println()
	pterm.DefaultParagraph.Println(
		pterm.Yellow(
			"This path will be available only during the execution of this program",
		),
	)
	if len(tfPath) == 0 {
		tfPath = os.Getenv("HOME") + "/bin"
	}

	if _, err := os.Stat(tfPath); os.IsNotExist(err) {
		os.Mkdir(tfPath, 0755)
	}

	installer := &releases.LatestVersion{
		Product:     product.Terraform,
		InstallDir:  tfPath,
		Constraints: constraint,
	}

	_, err := installer.Install(context.Background())
	if err != nil {
		pterm.Fatal.Println(err)
		os.Exit(1)
	}

	installer.Validate()

	SetTFPath(tfPath)

}

func IsTFInstalledAndCompatible(constraint string) {

	cmd, err := exec.Command("terraform", "version", "-json").Output()
	if err != nil {
		pterm.Error.Println(err)
		confirmed, _ := pterm.DefaultInteractiveConfirm.
			WithDefaultValue(true).
			Show("Terraform is not installed, do you want to install it now?")
		if !confirmed {
			os.Exit(1)
		} else {
			InstallTerraform(constraint)
		}
	} else {

		json.Unmarshal(cmd, &tfVersion)

		v, err := version.NewVersion(tfVersion.Version)
		if err != nil {
			pterm.Fatal.Println()
		}

		constraints, err := version.NewConstraint(constraint)
		if err != nil {
			pterm.Error.Println(err)
		}

		if !constraints.Check(v) {
			pterm.Println(pterm.Red("Terraform " + tfVersion.Version + " installed doesn't match with " + constraint + " constraint"))
			os.Exit(1)
		} else {
			pterm.Printfln(pterm.Green("Valid Terraform version found in the path: ", tfVersion.Version))
		}
	}

}
