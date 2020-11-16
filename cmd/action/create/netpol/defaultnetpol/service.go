package defaultnetpol

import (
	apiv1 "k8s.io/api/core/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/awscnfm/v12/pkg/key"
	"github.com/giantswarm/awscnfm/v12/pkg/project"
)

// nginxTestPodService will create a service pointing to nginx test pod
func nginxTestPodService() *apiv1.Service {
	svc := &apiv1.Service{
		ObjectMeta: apismetav1.ObjectMeta{
			Name:      key.NetPolNginxSvcName,
			Namespace: key.NetPolTestNamespaceName,
			Labels: map[string]string{
				key.LabelApp:       name,
				key.LabelManagedBy: project.Name(),
			},
		},
		Spec: apiv1.ServiceSpec{
			Type: apiv1.ServiceTypeClusterIP,
			Ports: []apiv1.ServicePort{
				{
					Name:     "http",
					Protocol: apiv1.ProtocolTCP,
					Port:     80,
				},
			},
			Selector: map[string]string{
				key.LabelApp: key.NetPolNginxTestPodAppLabel,
			},
		},
	}
	return svc
}
