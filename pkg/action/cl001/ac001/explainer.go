package ac001

import (
	"context"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/giantswarm/microerror"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	crs, err := newCRs("https://g8s.codename.eu-central-1.aws.gigantic.io:443")
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
