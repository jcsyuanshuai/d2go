package trees

import (
	"fmt"
	"testing"
)

var root = &TreeNode{
	Value: 1,
	Left: &TreeNode{
		Value: 2,
		Right: &TreeNode{
			Value: 4,
		},
	},
	Right: &TreeNode{
		Value: 3,
		Right: &TreeNode{
			Value: 5,
		},
	},
}

func TestInorderTraverseByStack(t *testing.T) {
	ret := InorderTraverseByStack(root)
	fmt.Printf("%v", ret)
}

func TestInorderTraverse(t *testing.T) {
	ret := InorderTraverse(root)
	fmt.Printf("%v", ret)
}
