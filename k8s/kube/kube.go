package kube

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func NewK8sClient() (*kubernetes.Clientset, error) {
	kubeconfig := flag.String(
		"kubeconfig",
		"k8s/.kube/config/admin.conf",
		"(optional) absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	return kubernetes.NewForConfig(config)
}

var clientset *kubernetes.Clientset

func init() {
	clientset, _ = NewK8sClient()
}

type ResourceType string

const (
	ResourceTypePod     ResourceType = "pod"
	ResourceTypeService ResourceType = "service"
)

func GetPodInterface(namespace string) corev1.PodInterface {
	if namespace == "" {
		namespace = "default"
	}
	return clientset.CoreV1().Pods(namespace)
}

func GetServiceInterface(namespace string) corev1.ServiceInterface {
	if namespace == "" {
		namespace = "default"
	}
	return clientset.CoreV1().Services(namespace)
}
