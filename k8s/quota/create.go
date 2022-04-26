package quota

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Create(clientset kubernetes.Clientset) (*corev1.ResourceQuota, error) {

	return clientset.CoreV1().ResourceQuotas("").Create(
		context.TODO(),
		&corev1.ResourceQuota{
			TypeMeta: metav1.TypeMeta{
				Kind:       "",
				APIVersion: "",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "",
			},
			Spec: corev1.ResourceQuotaSpec{
				Hard: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceLimitsCPU:              resource.MustParse(""),
					corev1.ResourceLimitsMemory:           resource.MustParse(""),
					corev1.ResourceLimitsEphemeralStorage: resource.MustParse(""),
					corev1.ResourceRequestsCPU:            resource.MustParse(""),
					corev1.ResourceRequestsMemory:         resource.MustParse(""),
					corev1.ResourcePods:                   resource.MustParse("4"),
					corev1.ResourceReplicationControllers: resource.MustParse("4"),
					corev1.ResourceSecrets:                resource.MustParse(""),
					corev1.ResourceServices:               resource.MustParse(""),
					corev1.ResourceQuotas:                 resource.MustParse(""),
				},
				Scopes:        nil,
				ScopeSelector: nil,
			},
			Status: corev1.ResourceQuotaStatus{},
		},
		metav1.CreateOptions{},
	)
}

func Delete(clientset kubernetes.Clientset) error {
	return clientset.CoreV1().ResourceQuotas("").Delete(
		context.TODO(),
		"",
		metav1.DeleteOptions{},
	)
}
