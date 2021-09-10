package tree

func levelOrder(root *TreeNode) [][]int32 {
	levels := make([][]int32, 0)
	if root == nil {
		return levels
	}

	q := []*TreeNode{root}

	for len(q) > 0 { // 替代while
		// dequeue and create level for the length
		// of the currently enqueued level
		level := make([]int32, 0)
		levelLen := len(q)
		// 每次循环树的一层，队列中存放的为当前层有哪些结点
		for i := 0; i < levelLen; i++ {
			// add nodes value to the level
			n := q[0]
			level = append(level, n.Val)

			if n.Left != nil {
				q = append(q, n.Left)
			}

			if n.Right != nil {
				q = append(q, n.Right)
			}

			// dequeue the node
			q = q[1:]
		}
		levels = append(levels, level)
	}
	return levels
}
