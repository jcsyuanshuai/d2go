package po

import (
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Model struct {
	ipod corev1.PodInterface
}

func New(clientset *kubernetes.Clientset, typename string, namespace string) *Model {
	return &Model{
		ipod: clientset.CoreV1().Pods(namespace),
	}
}
