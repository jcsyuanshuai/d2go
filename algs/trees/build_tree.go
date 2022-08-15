package trees

func buildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	ret := new(TreeNode)
	ret.Value = preorder[0]
	i := 0
	for ; i < len(preorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	n := len(inorder[:i])
	ret.Left = buildTree(preorder[1:n+1], inorder[:i])
	ret.Right = buildTree(preorder[n+1:], inorder[i+1:])
	return ret
}
