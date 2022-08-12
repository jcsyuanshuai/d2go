package arrays

import "github.com/jcsys/d2go/algs"

func MaxProfit(prices []int) int {
	n := len(prices)

	if n <= 1 {
		return 0
	}

	min, max := prices[0], 0
	for i := 1; i < n; i++ {
		max = algs.Max(max, prices[i]-min)
		min = algs.Min(min, prices[i])
	}
	return max
}
