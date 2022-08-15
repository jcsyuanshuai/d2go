package trees

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Value == p.Value || root.Value == q.Value {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	return root
}
