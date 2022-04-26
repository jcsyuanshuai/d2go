package main

import (
	"github.com/jcsys/d2go/k8s/kube"
	"github.com/jcsys/d2go/k8s/ns"
	"github.com/jcsys/d2go/k8s/po"
	"github.com/jcsys/d2go/k8s/svc"
)

func main() {
	clientset, _ := kube.NewK8sClient()
	ns.List(clientset)

	po.List(clientset)

	svc.List(clientset)

}
