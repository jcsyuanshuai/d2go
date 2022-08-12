package arrays

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	ret := make([][]int, 0)

	n := len(nums)
	if n < 3 {
		return ret
	}
	sort.Ints(nums)
	hashmap := make(map[int]int, 0)
	for idx, num := range nums {
		hashmap[num] = idx
	}

	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			return ret
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a, b := -1, -1
		for j := i + 1; j < n-1; j++ {
			x, y := nums[i], nums[j]
			target := -(x + y)
			if k, ok := hashmap[target]; ok && j < k {
				if a == j || b == k {
					continue
				}
				fmt.Printf("%d %d %d\n", i, j, k)
				ret = append(ret, []int{x, y, target})
				a, b = j, k
			}
		}
	}
	return ret
}

func ThreeSum2(nums []int) [][]int {
	ret := make([][]int, 0)

	n := len(nums)
	if n < 3 {
		return ret
	}
	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return ret
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			x, y, z := nums[i], nums[j], nums[k]
			sum := x + y + z
			if sum == 0 {
				ret = append(ret, []int{x, y, z})
				for j < k && y == nums[j+1] {
					j++
				}
				for j < k && z == nums[k-1] {
					k--
				}
				j++
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return ret
}
