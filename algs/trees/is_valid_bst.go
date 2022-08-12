package trees

import "math"

func isValidBST(root *TreeNode) bool {
	var validate func(tree *TreeNode, min, max int) bool

	validate = func(tree *TreeNode, min, max int) bool {
		if tree == nil {
			return true
		}
		if tree.Value <= min || tree.Value >= max {
			return false
		}
		return validate(tree.Left, min, tree.Value) &&
			validate(tree.Right, tree.Value, max)
	}

	return validate(root, math.MinInt, math.MaxInt)
}
