package tree

import "math"

var max = math.MinInt32

func maxPathSum(root *TreeNode) int32 {
	return maxPathSumIterator(root)
}

func maxPathSumIterator(root *TreeNode) int32 {
	if root == nil {
		return 0
	}
	leftSum := int32(math.Max(float64(maxPathSumIterator(root.Left)), 0.0))
	rightSum := int32(math.Max(float64(maxPathSumIterator(root.Right)), 0.0))
	pathSum := rightSum + leftSum + root.Val

	max = int(math.Max(float64(max), float64(pathSum)))

	return root.Val + int32(math.Max(float64(rightSum), float64(leftSum)))
}
