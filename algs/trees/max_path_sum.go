package trees

import "math"

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	var maxGain func(node *TreeNode) int

	maxGain = func(node *TreeNode) int {
		if root == nil {
			return 0
		}
		lgain := max(maxGain(root.Left), 0)
		rgain := max(maxGain(root.Right), 0)
		curr := root.Value + lgain + rgain
		maxSum = max(maxSum, curr)
		return node.Value + max(lgain, rgain)
	}

	maxGain(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
