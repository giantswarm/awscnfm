package ac008

import (
	"context"

	"github.com/giantswarm/microerror"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v12/pkg/label"
)

func (e *Executor) execute(ctx context.Context) error {
	{
		err := e.clients.ControlPlane.CtrlClient().DeleteAllOf(
			ctx,
			&apiv1alpha2.Cluster{},
			client.InNamespace(metav1.NamespaceDefault),
			client.MatchingLabels{label.Cluster: e.tenantCluster},
		)
		if errors.IsNotFound(err) {
			// fall through
		} else if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
