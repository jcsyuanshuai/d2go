package arrays

// FindMin Binary Search
// Suppose an array of length n sorted in ascending order
// is rotated between 1 and n times.
func FindMin(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)/2
		// 特殊情况，提前退出
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		// 特殊情况，提前退出
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		// 左半部分依次递增，最小值在右半部分
		// 至少一次滚动，不存在最小值在第一位
		if nums[mid] > nums[0] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
