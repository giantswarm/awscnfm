package defaultnetpol

import (
	apiv1 "k8s.io/api/core/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/awscnfm/v12/pkg/key"
)

// netPolTestNamespace will create a new namespace which will be used for testing the network policy
func netPolTestNamespace() *apiv1.Namespace {
	ns := &apiv1.Namespace{
		ObjectMeta: apismetav1.ObjectMeta{
			Name:      key.NetPolTestNamespaceName,
		},
	}
	return ns
}
