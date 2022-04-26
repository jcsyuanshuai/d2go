package ns

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Delete(clientset *kubernetes.Clientset) {
	err := clientset.CoreV1().Namespaces().Delete(
		context.TODO(),
		"",
		metav1.DeleteOptions{},
	)
	if err != nil {
		return
	}

}
