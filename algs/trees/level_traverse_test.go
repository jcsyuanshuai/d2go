package trees

import (
	"fmt"
	"testing"
)

func TestLevelTraverse2(t *testing.T) {
	tree := NewTreeNode([]int{3, 9, 20, 1 << 32, 1 << 32, 15, 7})
	ret := LevelTraverse(tree)
	fmt.Printf("%v\n", ret)
}
