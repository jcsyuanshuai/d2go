package po

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func List(clientset *kubernetes.Clientset) {

	pods, _ := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Printf("%s\n", pod.Name)
	}

	podInterface := clientset.CoreV1().Pods("default")
	podInterface.Get(context.TODO(), "", metav1.GetOptions{
		ResourceVersion: "",
	})

}
