package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserServiceImpl_LoginOrRegister(t *testing.T) {
	s := UserServiceImpl{}

	u, err := s.LoginOrRegister(18080901232, 8979)

	assert.NoError(t, err)
	assert.Equal(t, "test", u.Username)
}
