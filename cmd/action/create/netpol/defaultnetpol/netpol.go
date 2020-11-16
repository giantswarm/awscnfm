package defaultnetpol

import (
	networkingv1 "k8s.io/api/networking/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/awscnfm/v12/pkg/key"
	"github.com/giantswarm/awscnfm/v12/pkg/project"
)

// defaultNetworkPolicy will create a 'deny-from-all-namespaces' network policy
// https://github.com/ahmetb/kubernetes-network-policy-recipes/blob/master/04-deny-traffic-from-other-namespaces.md
func defaultNetworkPolicy() *networkingv1.NetworkPolicy {
	np := &networkingv1.NetworkPolicy{
		ObjectMeta: apismetav1.ObjectMeta{
			Name:      key.NetPolName,
			Namespace: key.NetPolTestNamespaceName,
			Labels: map[string]string{
				key.LabelManagedBy: project.Name(),
			},
		},
		Spec: networkingv1.NetworkPolicySpec{
			Ingress: []networkingv1.NetworkPolicyIngressRule{
				{
					From: []networkingv1.NetworkPolicyPeer{
						{
							PodSelector: &apismetav1.LabelSelector{},
						},
					},
				},
			},
			PodSelector: apismetav1.LabelSelector{
				MatchLabels: map[string]string{},
			},
		},
	}
	return np
}
