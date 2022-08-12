package binary

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)

	if totalLen%2 == 1 {
		return float64(getKthElement(nums1, nums2, (totalLen+1)/2))
	}
	leftK := getKthElement(nums1, nums2, totalLen/2)
	rightK := getKthElement(nums1, nums2, totalLen/2+1)

	return float64(leftK+rightK) / 2.0
}

func getKthElement(nums1, nums2 []int, k int) int {
	idx1, idx2 := 0, 0
	for {
		if idx1 == len(nums1) {
			return nums2[idx2+k-1]
		}

		if idx2 == len(nums2) {
			return nums1[idx1+k-1]
		}

		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}

		half := k / 2

		newIdx1 := min(idx1+half, len(nums1)) - 1
		newIdx2 := min(idx2+half, len(nums2)) - 1
		v1, v2 := nums1[newIdx1], nums2[newIdx2]

		// todo
		if v1 <= v2 {
			k -= newIdx1 - idx1 + 1
			idx1 = newIdx1 + 1
		} else {
			k -= newIdx2 - idx2 + 1
			idx2 = newIdx2 + 1
		}
	}
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
