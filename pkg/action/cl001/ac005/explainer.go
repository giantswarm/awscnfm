package ac005

import (
	"context"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/giantswarm/microerror"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	crs, err := newCRs(ctx, "al9qy", "https://g8s.codename.eu-central-1.aws.gigantic.io:443")
	if err != nil {
		return "", microerror.Mask(err)
	}

	machineDeploymentCR, err := yaml.Marshal(crs.MachineDeployment)
	if err != nil {
		return "", microerror.Mask(err)
	}

	awsMachineDeploymentCR, err := yaml.Marshal(crs.AWSMachineDeployment)
	if err != nil {
		return "", microerror.Mask(err)
	}

	s := `
Create a basic Node Pool with all its defaults. Note that some information of
the output below slightly differs when executing conformance tests on
different control planes.

---
` + strings.TrimSpace(string(machineDeploymentCR)) + `
---
` + strings.TrimSpace(string(awsMachineDeploymentCR)) + `
	`

	return s, nil
}
