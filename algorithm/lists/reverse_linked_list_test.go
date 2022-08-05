package lists

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	input := BuildListNode([]int{10, 8, 6, 4, 2})
	print("=== input === \n")
	PrintListNode(input)
	output := ReverseList(input)
	print("=== output === \n")
	PrintListNode(output)

	input2 := BuildListNode([]int{5, 4, 3, 2, 1})
	print("=== input2 === \n")
	PrintListNode(input2)
	output2 := ReverseList2(input2)
	print("=== output2 === \n")
	PrintListNode(output2)
}

func PrintListNode(node *ListNode) {
	if node == nil {
		return
	}
	sep := ""
	for p := node; p != nil; p = p.Next {
		fmt.Printf("%s%d", sep, p.Val)
		sep = "->"
	}
	fmt.Println()
}

func BuildListNode(arr []int) *ListNode {
	ret := &ListNode{Val: arr[0]}
	p := ret
	for i := 1; i < len(arr); i++ {
		p.Next = &ListNode{Val: arr[i]}
		p = p.Next
	}
	return ret
}
