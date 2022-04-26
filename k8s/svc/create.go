package svc

import (
	"context"
	"github.com/jcsys/d2go/k8s/kube"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Create() {
	kube.GetServiceInterface("").Create(
		context.TODO(),
		&corev1.Service{
			TypeMeta:   metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{},
			Spec: corev1.ServiceSpec{
				Ports:                         nil,
				Selector:                      nil,
				ClusterIP:                     "",
				ClusterIPs:                    nil,
				Type:                          "",
				ExternalIPs:                   nil,
				SessionAffinity:               "",
				LoadBalancerIP:                "",
				LoadBalancerSourceRanges:      nil,
				ExternalName:                  "",
				ExternalTrafficPolicy:         "",
				HealthCheckNodePort:           0,
				PublishNotReadyAddresses:      false,
				SessionAffinityConfig:         nil,
				IPFamilies:                    nil,
				IPFamilyPolicy:                nil,
				AllocateLoadBalancerNodePorts: nil,
				LoadBalancerClass:             nil,
				InternalTrafficPolicy:         nil,
			},
			Status: corev1.ServiceStatus{},
		},
		metav1.CreateOptions{
			TypeMeta:        metav1.TypeMeta{},
			DryRun:          nil,
			FieldManager:    "",
			FieldValidation: "",
		},
	)
}
