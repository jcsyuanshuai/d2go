package composite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitDept(t *testing.T) {
	assert.Equal(t, 20, InitDept().Count())
}
