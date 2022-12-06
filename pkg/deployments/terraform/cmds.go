package terraform

import (
	"context"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
)

func CompatibleVersions() []*releases.ExactVersion {
	constraint, _ := version.NewConstraint("~> 1.2")
	versions := &releases.Versions{
		Product:     product.Terraform,
		Constraints: constraint,
	}
	srcs, _ := versions.List(context.Background())
	exactVersions := []*releases.ExactVersion{}
	for _, src := range srcs {
		g, ok := src.(*releases.ExactVersion)
		if ok {
			exactVersions = append(exactVersions, g)
		}
	}
	return exactVersions
}

func LatestCompatibleVersion() *releases.ExactVersion {
	versions := CompatibleVersions()
	if len(versions) == 0 {
		return nil
	}
	return versions[len(versions)-1]
}

func InstallTerraform() {

	// constraint, _ := version.NewConstraint("~1.0")
	// versions := &releases.Versions{
	// 	Product:     product.Terraform,
	// 	Constraints: constraint,
	// }
	// versions.List(context.Background())
	// installer := &releases.ExactVersion{
	// 	Product: product.Terraform,
	// }
	// execPath, err := installer.Install(context.Background())
	// pterm.Fatal.PrintOnError(err)

	// installer.Validate()
}
