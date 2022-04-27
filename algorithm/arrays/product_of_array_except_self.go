package arrays

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = 1
	}

	left, right := 1, 1
	for i := 0; i < n; i++ {
		ans[i] *= left
		left *= nums[i]
		ans[n-1-i] *= right
		right *= nums[n-1-i]
	}
	return ans
}
