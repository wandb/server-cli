package deployments

import (
	"github.com/pterm/pterm"
)

func GetDeploymentStrategy() {
	pterm.DefaultSection.Println("Deployment Strategy")

	instance := GetInstance()

	dtype := instance.GetType()
	if dtype == "" {
		pterm.Println("We offer many different types of deployment strategies.")
		pterm.Println()

		pterm.Bold.Println(
			pterm.Blue("W&B Managed Dedicated Cloud"),
		)
		pterm.DefaultParagraph.Println(
			"A managed, dedicated deployment on W&B's" +
				"single-tenant infrastructure in your choice of cloud region. You can" +
				"provide use with your own bucket for extra security.",
		)
		pterm.Println()
		pterm.Bold.Println(
			pterm.Blue("W&B Managed Private Cloud"),
		)
		pterm.DefaultParagraph.Println(
			"A managed, dedicated deployment on your " +
				"infrastructure. You will need to grant our services account access " +
				"to your environment to provision resources. We will keep your version " +
				"up to date with the latest features and security fixes.",
		)
		pterm.Println()
		pterm.Bold.Println(
			pterm.Blue("Self-Managed Private Cloud"),
		)
		pterm.DefaultParagraph.Println(
			"Set up a production deployment on a private cloud in just a few steps " +
				"using terraform scripts provided by W&B. It will be your responsibility " +
				"to maintain this instance.",
		)
		pterm.Println()
		pterm.Bold.Println(
			pterm.Blue("Self-Managed Bare Metal"),
		)
		pterm.DefaultParagraph.Println(
			"W&B supports setting up a production server on most bare metal servers in " +
				"your on-premise data centers. We strongly advise against this type of deployment.",
		)
		pterm.Println()

		newType, _ := pterm.DefaultInteractiveSelect.
			WithDefaultText("Select deployment type").
			WithOptions([]string{
				string(ManagedDedicatedCloud),
				string(ManagedPrivateCloud),
				string(PrivateCloud),
				string(BareMetal),
			}).
			Show()
		dtype, _ = ParseType(newType)
		instance.SetType(newType)
	}

	platform := instance.GetPlatform()
	if platform == "" {
		platformOptions := []string{}
		switch dtype {
		case ManagedDedicatedCloud:
			fallthrough
		case ManagedPrivateCloud:
			fallthrough
		case PrivateCloud:
			platformOptions = append(
				platformOptions,
				string(AWS),
				string(GCP),
				string(Azure),
			)
		case BareMetal:
			platformOptions = append(platformOptions, string(Host), string(Kubernetes))
		}
		splatform, _ := pterm.DefaultInteractiveSelect.
			WithDefaultText("Select deployment platform").
			WithOptions(platformOptions).
			Show()
		instance.SetPlatform(splatform)
		platform, _ = ParsePlatform(splatform)
	}

	engine := instance.GetEngine()
	if engine == "" {
		engine = Terraform
		if platform == Host {
			engine = Docker
		}
		if platform == Kubernetes {
			engine = HelmChart
		}
		instance.SetEngine(string(engine))
	}
	instance.Write()

	pterm.Success.Println(string(dtype) + " > " + string(platform) + " > " + string(engine))
}
