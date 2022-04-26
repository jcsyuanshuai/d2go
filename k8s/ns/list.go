package ns

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func List(clientset *kubernetes.Clientset) {
	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, ns := range namespaceList.Items {
		fmt.Printf("%s\n", ns.Name)
	}
}
