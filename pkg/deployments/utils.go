package deployments

import (
	"io"
	"net/http"

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

	b, err := io.ReadAll(resp.Body)
	pterm.Fatal.PrintOnError()

	return string(b)
}
