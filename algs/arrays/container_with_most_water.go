package arrays

import "github.com/jcsys/d2go/algs"

func MaxArea(height []int) int {
	area := 0
	for i, _ := range height {
		for j := i + 1; j < len(height); j++ {
			area = algs.Max(area,
				algs.Min(height[i], height[j])*(j-i))
		}
	}
	return area
}

func MaxAreaByTowPointer(height []int) int {
	l, r := 0, len(height)-1
	area := 0
	for l < r {
		//area = algs.Max(area,
		//	algs.Min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			area = algs.Max(area, height[l]*(r-l))
			l += 1 // 移动高度小的那边指针
		} else {
			area = algs.Max(area, height[r]*(r-l))
			r -= 1
		}
	}
	return area
}
