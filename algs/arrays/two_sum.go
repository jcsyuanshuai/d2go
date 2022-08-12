package arrays

func TwoSum(nums []int, target int) []int {
	hashmap := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if idx, ok := hashmap[target-nums[i]]; ok {
			return []int{idx, i}
		} else {
			hashmap[nums[i]] = i
		}
	}
	return []int{}
}
