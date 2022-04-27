package arrays

import "sort"

func containsDuplicate(nums []int) bool {
	return solution1(nums)
}

func solution1(nums []int) bool {
	n := len(nums)
	hash := make(map[int]int, n)
	for _, num := range nums {
		hash[num] = 1
	}

	if len(hash) < n {
		return true
	} else {
		return false
	}
}

func solution2(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}
	return false
}
