package arrays

import "github.com/jcsys/d2go/algorithm"

func maxProfit(prices []int) int {
	n := len(prices)

	if n <= 1 {
		return 0
	}

	min, max := prices[0], 0
	for i := 1; i < n; i++ {
		max = algorithm.Max(max, prices[i]-min)
		min = algorithm.Min(min, prices[i])
	}
	return max
}
