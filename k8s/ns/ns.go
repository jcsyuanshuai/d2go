package ns

import (
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Client struct {
	client v1.NamespaceInterface
}
