package sorts

//func QuickSort(arr []int) []int {
//	n := len(arr)
//
//	if n <= 1 {
//		return arr
//	}
//
//	pivot := arr[rand.Intn(n)]
//
//	l, h, m := make([]int, 0, n), make([]int, 0, n), make([]int, 0, n)
//
//	for _, item := range arr {
//		switch {
//		case item < pivot:
//			l = append(l, item)
//		case item == pivot:
//			m = append(m, item)
//		case item > pivot:
//			h = append(h, item)
//		}
//	}
//
//	l = QuickSort(l)
//	h = QuickSort(h)
//
//	l = append(l, m...)
//	l = append(l, h...)
//
//	return l
//}

func QuickSort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	return sort(arr, 0, n-1)
}

func sort(arr []int, l, r int) []int {
	if l < r && l >= 0 && r > 0 {
		pivot := arr[l]
		index := l + 1
		for i := index; i < r; i++ {
			if arr[i] < pivot {
				arr[i], arr[index] = arr[index], arr[i]
				index++
			}
		}
		arr[l], arr[index-1] = arr[index-1], arr[l]

		p := index - 1
		sort(arr, l, p)
		sort(arr, p+1, r)
	}
	return arr
}
