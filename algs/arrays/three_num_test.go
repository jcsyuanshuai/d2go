package arrays

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	//arr := []int{-1, 0, 1, 2, -1, -4}
	//arr := []int{0, 0, 0, 0}
	arr := []int{-1, 0, 1, 0}
	ret := ThreeSum2(arr)
	fmt.Printf("%v", ret)
}
