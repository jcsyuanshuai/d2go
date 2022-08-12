package lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	// 节点为 nil 或者只有一个节点直接返回
	if head == nil || head.Next == nil {
		return head
	}
	// 每次递归返回原链中的最后一个节点
	ret := ReverseList(head.Next)
	head.Next.Next = head // 头节点置为尾节点
	head.Next = nil       // 头节点后继置为空

	return ret
}

func ReverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next       // 保存当前节点后继，防止交换当前节点和前继节点指针是断链
		curr.Next = prev        // 当前节点指针指向前继节点
		prev, curr = curr, next // 向后移动双指针
	}
	return prev
}
