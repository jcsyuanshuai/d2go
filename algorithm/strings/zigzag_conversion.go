package strings

func convert(s string, numRows int) string {
	n, r := len(s), numRows

	if r == 1 || r >= len(s) {
		return s
	}

	t := 2*r - 2
	ans := make([]byte, 0, n)
	for i := 0; i < r; i++ {
		for j := 0; j+i < n; j += t {
			ans = append(ans, s[j+i])
			// todo
			if i > 0 && i < r-1 && j+t-i < n {
				ans = append(ans, s[j+t-i])
			}
		}
	}
	return string(ans)
}
