package arrays

import "fmt"

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = 1
	}
	fmt.Printf("%v\n", ret)
	left, right := 1, 1
	for i := 0; i < n; i++ {
		ret[i] *= left
		left *= nums[i]
		ret[n-1-i] *= right
		right *= nums[n-1-i]
		fmt.Printf("%v\n", ret)
	}
	return ret
}
