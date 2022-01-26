package proxy

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserProxy_Login(t *testing.T) {
	p := NewUserProxy(&User{})

	err := p.Login("test", "111")

	require.Nil(t, err)
}
