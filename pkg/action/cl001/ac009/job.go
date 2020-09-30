package ac009

import (
	"fmt"

	batchapiv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/awscnfm/v12/pkg/project"
)

const (
	jobPriorityClass  = "giantswarm-critical"
	iamRoleAnnotation = "iam.amazonaws.com/role"
)

func jobNetworkPolicy() *networkingv1.NetworkPolicy {
	name := fmt.Sprintf("%s-kiam-test", project.Name())

	np := &networkingv1.NetworkPolicy{
		TypeMeta: apismetav1.TypeMeta{
			Kind:       "Job",
			APIVersion: batchapiv1.GroupName,
		},
		ObjectMeta: apismetav1.ObjectMeta{
			Name:      name,
			Namespace: kubeSystemNamespace,
			Labels: map[string]string{
				"app":        name,
				"created-by": project.Name(),
			},
		},

		Spec: networkingv1.NetworkPolicySpec{
			Egress: []networkingv1.NetworkPolicyEgressRule{
				{},
			},
			Ingress: []networkingv1.NetworkPolicyIngressRule{
				{},
			},
			PodSelector: apismetav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": name,
				},
			},
			PolicyTypes: []networkingv1.PolicyType{
				"Ingress",
				"Egress",
			},
		},
	}

	return np
}

// awsApiCallJob will will spawn a pod which will do simple AWS api call to test Kiam functionality
func awsApiCallJob(dockerRegistry string, awsRegion string, clusterID string) *batchapiv1.Job {
	activeDeadlineSeconds := int64(60)
	backOffLimit := int32(10)
	completions := int32(1)
	parallelism := int32(1)
	name := fmt.Sprintf("%s-kiam-test", project.Name())
	cpu := resource.MustParse("50m")
	memory := resource.MustParse("50Mi")

	// we will test using route53Manager role because its the only extra role that is created by default
	iamRole := fmt.Sprintf("%s-Route53Manager-Role", clusterID)

	j := &batchapiv1.Job{
		TypeMeta: apismetav1.TypeMeta{
			Kind:       "Job",
			APIVersion: batchapiv1.GroupName,
		},
		ObjectMeta: apismetav1.ObjectMeta{
			Name:      name,
			Namespace: kubeSystemNamespace,
			Labels: map[string]string{
				"app":        name,
				"created-by": project.Name(),
			},
		},
		Spec: batchapiv1.JobSpec{
			Parallelism:  &parallelism,
			Completions:  &completions,
			BackoffLimit: &backOffLimit,
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: apismetav1.ObjectMeta{
					Name:      name,
					Namespace: kubeSystemNamespace,
					Annotations: map[string]string{
						iamRoleAnnotation: iamRole,
					},
					Labels: map[string]string{
						"app":        name,
						"created-by": project.Name(),
					},
				},
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
							Env: []apiv1.EnvVar{
								{
									Name:  "AWS_DEFAULT_REGION",
									Value: awsRegion,
								},
							},
							Command: []string{
								"/bin/sh",
								"-c",
								"sleep 10s && echo 'trying to list domains via Route53-Manager role' && /usr/local/bin/aws route53 list-hosted-zones",
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
	return fmt.Sprintf("%s/giantswarm/awscli:2.0.24", dockerRegistry)
}
