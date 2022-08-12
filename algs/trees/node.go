package trees

import (
	"container/list"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{
		list: list.New(),
	}
}

func (s *Stack) Push(node *TreeNode) {
	s.list.PushBack(node)
}

func (s *Stack) Pop() *TreeNode {
	e := s.list.Back()
	s.list.Remove(e)
	return e.Value.(*TreeNode)
}

func (s *Stack) Empty() bool {
	return s.list.Len() <= 0
}

func NewTreeNode(arr []int) *TreeNode {
	n := len(arr)
	if n == 0 {
		return nil
	}
	ret := make([]*TreeNode, n)

	var p *TreeNode
	for i := 0; i < n; i++ {
		if arr[i] != 1<<32 {
			p = &TreeNode{Value: arr[i]}
		} else {
			p = nil
		}
		ret[i] = p
		if i > 0 {
			idx := (i - 1) >> 1
			if (i % 2) != 0 {
				ret[idx].Left = p
			} else {
				ret[idx].Right = p
			}
		}
	}
	return ret[0]
}
