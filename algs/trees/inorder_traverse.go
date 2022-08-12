package trees

func InorderTraverse(root *TreeNode) []int {
	ret := make([]int, 0)

	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		ret = append(ret, node.Value)
		traverse(node.Right)

	}
	traverse(root)

	return ret
}

func InorderTraverseByStack(root *TreeNode) []int {
	ret := make([]int, 0)
	s, p := NewStack(), root
	for p != nil || !s.Empty() {
		for p != nil {
			s.Push(p)
			p = p.Left
		}

		if !s.Empty() {
			p = s.Pop()
			ret = append(ret, p.Value)
			p = p.Right
		}
	}
	return ret
}
