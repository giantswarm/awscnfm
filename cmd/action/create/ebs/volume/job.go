package volume

import (
	"fmt"

	batchapiv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/awscnfm/v15/pkg/project"
)

const (
	jobPriorityClass = "giantswarm-critical"
)

func ensurePersistentVolumeClaim() *apiv1.PersistentVolumeClaim {
	return &apiv1.PersistentVolumeClaim{
		TypeMeta: apismetav1.TypeMeta{
			Kind:       "PersistentVolumeClaim",
			APIVersion: apiv1.GroupName,
		},
		ObjectMeta: apismetav1.ObjectMeta{
			Name:      "ebs-claim",
			Namespace: "default",
		},
		Spec: apiv1.PersistentVolumeClaimSpec{
			AccessModes: []apiv1.PersistentVolumeAccessMode{
				apiv1.ReadWriteOnce,
			},
			Resources: apiv1.ResourceRequirements{
				Requests: apiv1.ResourceList{
					apiv1.ResourceName(v1.ResourceStorage): resource.MustParse("5Gi"),
				},
			},
		},
	}

}

func ensureClusterRole() *rbacv1.ClusterRole {
	return &rbacv1.ClusterRole{
		TypeMeta: apismetav1.TypeMeta{
			Kind:       "ClusterRole",
			APIVersion: rbacv1.GroupName,
		},
		ObjectMeta: apismetav1.ObjectMeta{
			Name: "enable-ebs-psp",
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups:     []string{"policy"},
				Resources:     []string{"podsecuritypolicies"},
				ResourceNames: []string{"privileged"},
				Verbs:         []string{"use"},
			},
		},
	}
}

func ensureRoleBinding() *rbacv1.RoleBinding {
	return &rbacv1.RoleBinding{
		TypeMeta: apismetav1.TypeMeta{
			Kind:       "RoleBinding",
			APIVersion: rbacv1.GroupName,
		},
		ObjectMeta: apismetav1.ObjectMeta{
			Name:      "ebs-rolebinding",
			Namespace: "default",
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacv1.GroupName,
			Kind:     "ClusterRole",
			Name:     "enable-ebs-psp",
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      "default",
				Namespace: "default",
			},
		},
	}
}

// ensureEBSVolume will will spawn a pod which to ensure EBS volume attachment works.
func ensureEBSVolume(dockerRegistry string, clusterID string) *batchapiv1.Job {
	backOffLimit := int32(20)
	completions := int32(1)
	parallelism := int32(1)
	name := fmt.Sprintf("%s-ebs-volume-test", project.Name())
	cpu := resource.MustParse("50m")
	memory := resource.MustParse("50Mi")

	job := &batchapiv1.Job{
		TypeMeta: apismetav1.TypeMeta{
			Kind:       "Job",
			APIVersion: batchapiv1.GroupName,
		},
		ObjectMeta: apismetav1.ObjectMeta{
			Name:      name,
			Namespace: "default",
			Labels: map[string]string{
				"app":        name,
				"managed-by": project.Name(),
			},
		},
		Spec: batchapiv1.JobSpec{
			Parallelism:  &parallelism,
			Completions:  &completions,
			BackoffLimit: &backOffLimit,
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: apismetav1.ObjectMeta{
					Name:      name,
					Namespace: "default",
					Labels: map[string]string{
						"app":        name,
						"managed-by": project.Name(),
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
							Command: []string{
								"/bin/sh",
								"-c",
								"sleep 10s && echo $(date -u) >> /data/out.txt",
							},
							VolumeMounts: []apiv1.VolumeMount{
								{
									Name:      "persistent-storage",
									MountPath: "/data",
								},
							},
						},
					},
					RestartPolicy:     apiv1.RestartPolicyNever,
					PriorityClassName: jobPriorityClass,
					Volumes: []apiv1.Volume{
						{
							Name: "persistent-storage",
							VolumeSource: apiv1.VolumeSource{
								PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
									ClaimName: "ebs-claim",
								},
							},
						},
					},
				},
			},
		},
	}

	return job
}

func jobDockerImage(dockerRegistry string) string {
	return fmt.Sprintf("%s/giantswarm/awscli:2.0.24", dockerRegistry)
}
