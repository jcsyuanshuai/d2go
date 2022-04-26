package binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	ret := FindMedianSortedArrays([]int{1, 2}, []int{3, 4})
	assert.Equal(t, 2.5, ret)
}
