package label

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIsMaster(t *testing.T) {
	testCases := []struct {
		name     string
		node     corev1.Node
		expected bool
	}{
		{
			name: "case 0: kubeadm version >=1.25",
			node: corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"node-role.kubernetes.io/control-plane": "",
					},
				},
			},
			expected: true,
		},
		{
			name: "case 1: kubeadm version <1.25",
			node: corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"node-role.kubernetes.io/master": "",
					},
				},
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsMaster(tc.node)
			if actual != tc.expected {
				t.Fatalf("expected %v to be equal to %v", actual, tc.expected)
			}
		})
	}

}
