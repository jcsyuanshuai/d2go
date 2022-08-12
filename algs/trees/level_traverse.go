package trees

import (
	"container/list"
)

func LevelTraverse(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		tmp := make([]int, 0)
		l := q.Len()
		for i := 0; i < l; i++ {
			f := q.Front()
			n := f.Value.(*TreeNode)
			tmp = append(tmp, n.Value)
			if n.Left != nil {
				q.PushBack(n.Left)
			}
			if n.Right != nil {
				q.PushBack(n.Right)
			}
			q.Remove(f)
		}
		ret = append(ret, tmp)
	}
	return ret
}
