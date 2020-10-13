package ac008

import (
	"context"
	"net"

	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pkgclient "github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/label"
)

func (e *Executor) execute(ctx context.Context) error {
	var err error

	var cpClients k8sclient.Interface
	{
		c := pkgclient.ControlPlaneConfig{
			Logger: e.logger,

			KubeConfig: env.ControlPlaneKubeConfig(),
		}

		cpClients, err = pkgclient.NewControlPlane(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	{
		var list infrastructurev1alpha2.AWSClusterList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: e.tenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}

		for _, cluster := range list.Items {
			subnetString := cluster.Status.Provider.Network.CIDR

			_, npSubnet, err := net.ParseCIDR("192.168.64.0/18")
			if err != nil {
				return microerror.Mask(err)
			}

			cpIPs, _, err := net.ParseCIDR(subnetString)
			if err != nil {
				return microerror.Mask(err)
			}

			// This only checks if the IP part of the CIDR is inside the NetworkPool CIDR
			if !npSubnet.Contains(cpIPs) {
				return microerror.Mask(customNetworkPoolMasterUnusedError)
			}
		}
	}

	{
		var list infrastructurev1alpha2.AWSMachineDeploymentList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: e.tenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}

		for _, md := range list.Items {
			subnetString := md.Annotations["machine-deployment.giantswarm.io/subnet"]

			_, npSubnet, err := net.ParseCIDR("192.168.64.0/18")
			if err != nil {
				return microerror.Mask(err)
			}

			mdIP, _, err := net.ParseCIDR(subnetString)
			if err != nil {
				return microerror.Mask(err)
			}

			// This only checks if the IP part of the CIDR is inside the NetworkPool CIDR
			if !npSubnet.Contains(mdIP) {
				return microerror.Mask(customNetworkPoolWorkerUnusedError)
			}
		}
	}

	return nil
}
