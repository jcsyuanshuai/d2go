package sorts

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 3, 2, 4, 1}
	arr1 := BubbleSort(arr)
	fmt.Printf("%v", arr1)
}
