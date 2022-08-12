package algs

func Min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}
