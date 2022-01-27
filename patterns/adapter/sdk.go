package adapter

import "fmt"

type AwsClient struct {
}

func (a *AwsClient) Run(cpu, mem float64) error {
	fmt.Printf("aws client run success, cpu: %f, mem: %f,", cpu, mem)
	return nil
}

type AliClient struct {
}

func (a *AliClient) Run(cpu, mem int) error {
	fmt.Printf("ali client run success, cpu: %d, mem: %d,", cpu, mem)
	return nil
}

type Adaptor interface {
	CreateServer(cpu, mem float64) error
}

type AwsClientAdapter struct {
	client AwsClient
}

func (a *AwsClientAdapter) CreateServer(cpu, mem float64) error {
	return a.client.Run(cpu, mem)
}

type AliClientAdapter struct {
	client AliClient
}

func (a *AliClientAdapter) CreateServer(cpu, mem float64) error {
	return a.client.Run(int(cpu), int(mem))
}
