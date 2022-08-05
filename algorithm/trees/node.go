package trees

import "container/list"

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
