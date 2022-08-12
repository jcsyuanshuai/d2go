package lru

type LinkedNode struct {
	Value pair
	Prev  *LinkedNode
	Next  *LinkedNode
	list  *LinkedList
}

type LinkedList struct {
	Len  int
	root LinkedNode
}

func NewLinkedList() *LinkedList {
	l := &LinkedList{
		Len: 0,
	}
	l.root.Next = &l.root
	l.root.Prev = &l.root
	return l
}

func (l *LinkedList) MoveToFront(e *LinkedNode) {
	if e.list != l || l.root.Next == e {
		return
	}
	l.move(e, &l.root)
}

func (l *LinkedList) move(e, at *LinkedNode) {
	if e == at {
		return
	}
	// 移动节点前继指向后继
	e.Prev.Next = e.Next
	// 移动节点后继指向前继
	e.Next.Prev = e.Prev

	// 移动节点前继指向目标节点
	e.Prev = at
	// 移动节点后继指向目标节点的后继
	e.Next = at.Next
	// 目标节点后继指向移动节点
	e.Prev.Next = e
	// 目标节点后继指向移动节点
	e.Next.Prev = e
}

func (l *LinkedList) insert(e, at *LinkedNode) {
	e.Prev = at
	e.Next = at.Next
	e.Prev.Next = e
	e.Next.Prev = e
	e.list = l
	l.Len++
}

func (l *LinkedList) PushToFront(p pair) {
	l.insert(&LinkedNode{Value: p}, &l.root)
}

func (l *LinkedList) Front() *LinkedNode {
	if l.Len == 0 {
		return nil
	}
	return l.root.Next
}

func (l *LinkedList) Back() *LinkedNode {
	if l.Len == 0 {
		return nil
	}
	return l.root.Prev
}

func (l *LinkedList) Remove(e *LinkedNode) pair {
	if e.list != l {
		return pair{}
	}
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev
	e.Next = nil
	e.Prev = nil
	e.list = nil
	l.Len--
	return e.Value
}
