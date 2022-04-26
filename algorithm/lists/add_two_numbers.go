package lists

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := ListNode{}
	cursor := &root
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		cursor.Next = &ListNode{
			Val: sum % 10,
		}

		cursor = cursor.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return root.Next
}
