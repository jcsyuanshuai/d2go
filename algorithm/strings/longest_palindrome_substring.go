package strings

func longestPalindrome1(s string) string {
	l := len(s)

	if l < 2 {
		return s
	}

	start := 0
	maxL := 0
	dp := make([][]bool, l, l)
	for i := 0; i < l; i++ {
		dp[i][i] = true
	}

	for subL := 2; subL <= l; subL++ {
		for i := 0; i < l; i++ {
			j := subL + i - 1
			if j >= l {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxL {
				maxL = j - i + 1
				start = i
			}
		}
	}

	return s[start : start+maxL]
}

//  abccbaf
//  abbcbbad

func longestPalindrome(s string) string {

	if len(s) < 2 || len(s) == 2 && s[0] == s[1] {
		return s
	}

	idx1, idx2 := 0, 0
	for mid := 1; mid < len(s)-1; mid++ {
		if s[mid-1] == s[mid] {
			idx1, idx2 = center(mid-1, mid, idx1, idx2, s)
		}

		if s[mid] == s[mid+1] {
			idx1, idx2 = center(mid, mid+1, idx1, idx2, s)
		}

		if s[mid-1] == s[mid+1] {
			idx1, idx2 = center(mid-1, mid+1, idx1, idx2, s)
		}
	}
	return s[idx1 : idx2+1]
}

func center(i, j, idx1, idx2 int, s string) (int, int) {
	for i > 0 && j < len(s)-1 {
		if s[i-1] == s[j+1] {
			i--
			j++
		} else {
			break
		}
	}

	if j-i+1 > idx2-idx1+1 {
		return i, j
	} else {
		return idx1, idx2
	}
}
