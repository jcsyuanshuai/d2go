package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func compute(factory Factory, x, y int) int {
	operator := factory.Create()
	operator.SetX(x)
	operator.SetY(y)
	return operator.Result()
}

func TestMinusFactory_Create(t *testing.T) {
	assert.Equal(t, 5, compute(MinusFactory{}, 10, 5))
}

func TestPlusFactory_Create(t *testing.T) {
	assert.Equal(t, 5, compute(PlusFactory{}, 1, 4))
}
