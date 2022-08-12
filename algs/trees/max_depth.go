package trees

import "github.com/jcsys/d2go/algs"

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := MaxDepth(root.Left)
	r := MaxDepth(root.Right)
	return algs.Max(l, r) + 1
}
