package strings

// 滑动窗口 abcabcabb
// start 0  0   0   0   1   2   3
// ret   0  1   2   3   3   3   3

func LengthOfLongestSubstring(s string) int {
	last := make([]int, 128)
	for i := 0; i < 128; i++ {
		last[i] = -1
	}

	start := 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		idx := s[i]
		if start < last[idx]+1 {
			start = last[idx] + 1
		}

		if maxLen < i-start+1 {
			maxLen = i - start + 1
		}

		last[idx] = i
	}
	return maxLen
}
