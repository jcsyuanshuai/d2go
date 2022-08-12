package arrays

func PivotIndex(nums []int) int {
	for index, _ := range nums {
		sum1 := 0
		for i := 0; i < index; i++ {
			sum1 += nums[i]
		}
		sum2 := 0
		for j := len(nums) - 1; j > index; j-- {
			sum2 += nums[j]
		}
		if sum1 == sum2 {
			return index
		}
	}
	return -1
}
