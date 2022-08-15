package trees

import (
	"container/list"
)

// 最优
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	l := invertTree(root.Left)
	r := invertTree(root.Right)
	root.Left = r
	root.Right = l
	return root
}

func invertTreeByQueue(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		e := q.Front().Value.(*TreeNode)
		e.Left, e.Right = e.Right, e.Left
		if e.Left != nil {
			q.PushBack(e.Left)
		}
		if e.Right != nil {
			q.PushBack(e.Right)
		}
	}
	return root
}

func invertTreeByStack(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	s := list.New()
	s.PushBack(root)
	for s.Len() > 0 {
		size := s.Len()
		for i := 0; i < size; i++ {
			e := s.Back().Value.(*TreeNode)
			e.Left, e.Right = e.Right, e.Left
			if e.Right != nil {
				s.PushBack(e.Right)
			}
			if e.Left != nil {
				s.PushBack(e.Left)
			}
		}
	}
	return root
}
