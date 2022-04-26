package ns

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Update(clientset *kubernetes.Clientset) {
	clientset.CoreV1().Namespaces().Update(
		context.TODO(),
		&corev1.Namespace{
			TypeMeta:   metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{},
			Spec:       corev1.NamespaceSpec{},
			Status:     corev1.NamespaceStatus{},
		},
		metav1.UpdateOptions{},
	)
}
