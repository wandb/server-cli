package deployments

import (
	"io"
	"net/http"

	"github.com/pterm/pterm"
)

func GetTerraformTemplate(name string) string {
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
	pterm.Fatal.PrintOnError(err)

	client := &http.Client{}
	resp, err := client.Do(req)
	pterm.Fatal.PrintOnError(err)

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	pterm.Fatal.PrintOnError()

	return string(b)
}
