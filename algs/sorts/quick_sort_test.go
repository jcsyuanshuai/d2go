package sorts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1, 3, 5, 9, 2, 6, 7, 4, 8}
	ret := QuickSort(arr)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, ret)
}
