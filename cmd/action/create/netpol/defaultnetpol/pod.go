package defaultnetpol

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/awscnfm/v12/pkg/key"
	"github.com/giantswarm/awscnfm/v12/pkg/project"
)

// nginxTestPod will create a test pod running nginx in the test namespace
func nginxTestPod(dockerRegistry string) *apiv1.Pod {
	cpu := resource.MustParse("250m")
	memory := resource.MustParse("200Mi")

	pod := &apiv1.Pod{
		ObjectMeta: apismetav1.ObjectMeta{
			Name:      key.NetPolNginxTestPodName,
			Namespace: key.NetPolTestNamespaceName,
			Labels: map[string]string{
				key.LabelApp:       key.NetPolNginxTestPodAppLabel,
				key.LabelManagedBy: project.Name(),
			},
		},
		Spec: apiv1.PodSpec{
			Containers: []apiv1.Container{
				{
					Name:  name,
					Image: podDockerImage(dockerRegistry),
					Resources: apiv1.ResourceRequirements{
						Limits: apiv1.ResourceList{
							apiv1.ResourceCPU:    cpu,
							apiv1.ResourceMemory: memory,
						},
						Requests: apiv1.ResourceList{
							apiv1.ResourceCPU:    cpu,
							apiv1.ResourceMemory: memory,
						},
					},
					Ports: []apiv1.ContainerPort{
						{
							ContainerPort: 80,
							Name:          "http",
							Protocol:      apiv1.ProtocolTCP,
						},
					},
				},
			},
			RestartPolicy: apiv1.RestartPolicyNever,
		},
	}

	return pod
}

func podDockerImage(dockerRegistry string) string {
	return fmt.Sprintf("%s/giantswarm/nginx:1.18-alpine", dockerRegistry)
}
