package ac001

import (
	"context"
	"strings"

	"github.com/giantswarm/awscnfm/pkg/env"
	"github.com/giantswarm/microerror"
	"gopkg.in/yaml.v2"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	server, err := env.CurrentServer(env.KubeConfig())
	if err != nil {
		return "", microerror.Mask(err)
	}
	crs, npCrs, err := newCRs(server)
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

	machineDeploymentCR, err := yaml.Marshal(npCrs.MachineDeployment)
	if err != nil {
		return "", microerror.Mask(err)
	}

	awsMachineDeploymentCR, err := yaml.Marshal(npCrs.AWSMachineDeployment)
	if err != nil {
		return "", microerror.Mask(err)
	}

	s := `
Create a basic Tenant Cluster with all its defaults and one additional node pool. Note that some
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
---
` + strings.TrimSpace(string(machineDeploymentCR)) + `
---
` + strings.TrimSpace(string(awsMachineDeploymentCR)) + `
	`

	return s, nil
}
