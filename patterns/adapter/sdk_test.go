package adapter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAliClientAdapter_CreateServer(t *testing.T) {
	a := &AliClientAdapter{
		client: AliClient{},
	}

	err := a.CreateServer(1.0, 2.0)

	require.Nil(t, err)
}

func TestAwsClientAdapter_CreateServer(t *testing.T) {
	a := &AwsClientAdapter{
		client: AwsClient{},
	}

	err := a.CreateServer(1.0, 2.0)

	require.Nil(t, err)
}
