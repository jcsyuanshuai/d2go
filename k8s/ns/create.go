package ns

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func Create(clientset *kubernetes.Clientset) {
	_, err := clientset.CoreV1().Namespaces().Create(
		context.TODO(),
		&corev1.Namespace{
			TypeMeta: metav1.TypeMeta{
				Kind:       "",
				APIVersion: "",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "test",
			},
			Spec: corev1.NamespaceSpec{
				Finalizers: nil,
			},
			Status: corev1.NamespaceStatus{
				Phase:      "",
				Conditions: nil,
			},
		},
		metav1.CreateOptions{},
	)
	if err != nil {
		return
	}
}
