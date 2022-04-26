package svc

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func List(clientset *kubernetes.Clientset) {
	services, err := clientset.CoreV1().Services("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return
	}
	for _, item := range services.Items {
		fmt.Printf("%s\n", item.Name)
	}

}
