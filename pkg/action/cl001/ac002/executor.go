package ac002

import (
	"context"

	"github.com/giantswarm/microerror"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (e *Executor) execute(ctx context.Context) error {
	_, err := e.clients.TenantCluster.K8sClient().CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
