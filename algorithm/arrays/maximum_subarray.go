package arrays

import "github.com/jcsys/d2go/algorithm"

func maxSubArray(nums []int) int {
	ans, sum := nums[0], 0

	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		ans = algorithm.Max(ans, sum)
	}
	return ans
}

func solution(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = algorithm.Max(nums[i-1]+nums[i], nums[i])
		ans = algorithm.Max(ans, nums[i])
	}
	return ans
}
