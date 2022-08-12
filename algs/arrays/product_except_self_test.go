package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	arr := []int{-1, 1, 0, -3, 3}

	ret := ProductExceptSelf(arr)

	assert.Equal(t, []int{0, 0, 9, 0, 0}, ret)
}
