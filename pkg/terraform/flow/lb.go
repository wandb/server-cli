package flow

import (
	"net/netip"
	"strings"

	"github.com/pterm/pterm"
	"github.com/wandb/server-cli/pkg/terraform/tfconfig"
)

type LoadBalancerConfig struct {
}

func AllowListFlow(config *tfconfig.TerraformConfig) {
	pterm.DefaultParagraph.
		Println(
			"An allow list lets you set which IP address can connect " +
				"to your instance. Default is all IP addresses (0.0.0.0/0 and ::/0). ",
		)

	enable, _ := pterm.
		DefaultInteractiveConfirm.
		Show("Would you like to configure an allow list?")
	if !enable {
		return
	}

	areCidrsValid := true
	for {
		cidrs, _ := pterm.DefaultInteractiveTextInput.
			WithMultiLine().
			Show("Please put each CIDR range on a new line")
		strings.Split(cidrs, "\n")

		validCidrs6 := []string{}
		validCidrs4 := []string{}
		for _, cidr := range strings.Split(cidrs, "\n") {
			if cidr == "" {
				continue
			}

			cidrObj, _ := netip.ParsePrefix(strings.TrimSpace(cidr))
			if !cidrObj.IsValid() {
				areCidrsValid = false
				pterm.Warning.Printfln("\"%s\" is not a valid CIDR", cidr)
			}

			if cidrObj.Addr().Is6() {
				validCidrs6 = append(validCidrs6, cidr)
			}
			if cidrObj.Addr().Is4() {
				validCidrs4 = append(validCidrs4, cidr)
			}
		}

		if areCidrsValid {
			return
		}
	}
}
