package trees

import "container/list"

func LevelTraverse(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	q := list.New()
	q.PushBack(root)

	for q.Len() != 0 {
		count := 1
		tmp := make([]int, 0)
		for i := 1; i <= count; i++ {
			f := q.Front()
			n := f.Value.(TreeNode)
			if n.Left != nil {
				q.PushBack(n.Left)
				count++
			}
			if n.Right != nil {
				q.PushBack(n.Right)
				count++
			}
			q.Remove(f)
		}
		count--
		ret = append(ret, tmp)
	}
	return ret
}
