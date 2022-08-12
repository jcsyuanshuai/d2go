package arrays

import "github.com/jcsys/d2go/algs"

// MaxProduct
// A subarray is a contiguous subsequence of the array.
func MaxProduct(nums []int) int {
	ret := nums[0]
	for i, _ := range nums {
		mul := nums[i]
		for j := i + 1; j < len(nums); j++ {
			ret = algs.Max(ret, mul)
			mul *= nums[j]
		}
		ret = algs.Max(ret, mul)
	}
	return ret
}

func MaxProductByTwoPointer(nums []int) int {
	// todo
	ret := 0
	return ret
}
