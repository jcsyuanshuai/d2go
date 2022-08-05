package sorts

func MergeSort(arr []int) []int {
	n := len(arr)

	if n < 2 {
		return arr
	}

	m := n / 2

	x := MergeSort(arr[:m])
	y := MergeSort(arr[m:])

	return merge(x, y)
}

func merge(x, y []int) []int {
	n1, n2 := len(x), len(y)
	ret := make([]int, n1+n2)

	i, j := 0, 0
	for i < n1 && j < n2 {
		if x[i] <= y[j] {
			ret[i+j] = x[i]
			i++
		} else {
			ret[i+j] = y[j]
			j++
		}
	}
	for i < n1 {
		ret[i+j] = x[i]
		i++
	}

	for j < n2 {
		ret[i+j] = y[j]
		j++
	}

	return ret
}
