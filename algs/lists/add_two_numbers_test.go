package lists

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	addTwoNumbers(l1, l2)
}
