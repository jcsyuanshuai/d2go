package arrays

import "github.com/jcsys/d2go/algs"

func MaxSubArray(nums []int) int {
	ans, sum := nums[0], 0

	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		ans = algs.Max(ans, sum)
	}
	return ans
}

func solution(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = algs.Max(nums[i-1]+nums[i], nums[i])
		ans = algs.Max(ans, nums[i])
	}
	return ans
}
