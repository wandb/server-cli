package deployments

import (
	"io"
	"net/http"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/pterm/pterm"
)

func GetTerraformTemplate(name string) string {
	url := "https://raw.githubusercontent.com/wandb/server-cli/main/templates/" + name + ".tf"
	req, err := http.NewRequest("GET", url, nil)
	pterm.Fatal.PrintOnError(err)

	client := &http.Client{}
	resp, err := client.Do(req)
	pterm.Fatal.PrintOnError(err)

	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)
	pterm.Fatal.PrintOnError()

	return string(b)
}

func PetName() string {
	petName := petname.Generate(2, "-")
	return petName
}
