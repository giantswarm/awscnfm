package ac001

import (
	"context"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger/microloggertest"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	var err error

	var cpClients k8sclient.Interface
	{
		c := client.ControlPlaneConfig{
			Logger: microloggertest.New(),

			KubeConfig: env.KubeConfig(),
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

	crs, npcrs, err := newCRs(releases, "https://g8s.codename.eu-central-1.aws.gigantic.io:443")
	if err != nil {
		return "", microerror.Mask(err)
	}

	networkPoolCR, err := yaml.Marshal(npcrs.NetworkPool)
	if err != nil {
		return "", microerror.Mask(err)
	}

	clusterCR, err := yaml.Marshal(crs.Cluster)
	if err != nil {
		return "", microerror.Mask(err)
	}

	awsClusterCR, err := yaml.Marshal(crs.AWSCluster)
	if err != nil {
		return "", microerror.Mask(err)
	}

	g8sControlPlaneCR, err := yaml.Marshal(crs.G8sControlPlane)
	if err != nil {
		return "", microerror.Mask(err)
	}

	awsControlPlaneCR, err := yaml.Marshal(crs.AWSControlPlane)
	if err != nil {
		return "", microerror.Mask(err)
	}

	s := `
Create a basic Tenant Cluster with all its defaults. Note that some
information of the output below slightly differs when executing conformance
tests on different control planes.

---
` + strings.TrimSpace(string(networkPoolCR)) + `
---
` + strings.TrimSpace(string(clusterCR)) + `
---
` + strings.TrimSpace(string(awsClusterCR)) + `
---
` + strings.TrimSpace(string(g8sControlPlaneCR)) + `
---
` + strings.TrimSpace(string(awsControlPlaneCR)) + `
	`

	return s, nil
}
