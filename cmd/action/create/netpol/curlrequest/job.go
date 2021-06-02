package curlrequest

import (
	"fmt"

	batchapiv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/awscnfm/v14/pkg/key"
	"github.com/giantswarm/awscnfm/v14/pkg/project"
)

const (
	jobPriorityClass = "giantswarm-critical"
)

// testNetworkPolicyJob will spawn a pod which will do curl request to nginx test service to test network connectivity
func testNetworkPolicyJob(dockerRegistry string, namespace string) *batchapiv1.Job {
	activeDeadlineSeconds := int64(600)
	backOffLimit := int32(5)
	completions := int32(1)
	parallelism := int32(1)
	cpu := resource.MustParse("50m")
	memory := resource.MustParse("50Mi")

	jobName := key.NetPolTestJobName(namespace)

	j := &batchapiv1.Job{
		TypeMeta: apismetav1.TypeMeta{
			Kind:       "Job",
			APIVersion: batchapiv1.GroupName,
		},
		ObjectMeta: apismetav1.ObjectMeta{
			Name:      jobName,
			Namespace: namespace,
			Labels: map[string]string{
				key.LabelManagedBy: project.Name(),
			},
		},
		Spec: batchapiv1.JobSpec{
			Parallelism:  &parallelism,
			Completions:  &completions,
			BackoffLimit: &backOffLimit,
			Template: apiv1.PodTemplateSpec{
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  name,
							Image: jobDockerImage(dockerRegistry),
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
							Command: []string{
								"/bin/sh",
								"-c",
								key.NetPolCurlTestCommand(),
							},
						},
					},
					RestartPolicy:     apiv1.RestartPolicyNever,
					PriorityClassName: jobPriorityClass,
				},
			},
			ActiveDeadlineSeconds: &activeDeadlineSeconds,
		},
	}

	return j
}

func jobDockerImage(dockerRegistry string) string {
	return fmt.Sprintf("%s/giantswarm/curl:7.73.0", dockerRegistry)
}
