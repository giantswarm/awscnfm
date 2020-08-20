package ac005

import (
	"context"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger/microloggertest"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	var err error

	var cpClients k8sclient.Interface
	{
		c := client.ControlPlaneConfig{
			Logger: microloggertest.New(),
		}

		cpClients, err = client.NewControlPlane(c)
		if err != nil {
			return "", microerror.Mask(err)
		}
	}

	var releases []v1alpha1.Release
	{
		var list v1alpha1.ReleaseList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
		)
		if err != nil {
			return "", microerror.Mask(err)
		}

		releases = list.Items
	}

	crs, err := newCRs(ctx, releases, "al9qy", "https://g8s.codename.eu-central-1.aws.gigantic.io:443")
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
