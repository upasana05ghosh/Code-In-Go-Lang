//https://leetcode.com/problems/binary-tree-maximum-path-sum/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var max int

func maxPathSum(root *TreeNode) int {
	max = math.MinInt32
	findMax(root)
	return max
}

func findMax(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := calMax(0, findMax(root.Left))
	rh := calMax(0, findMax(root.Right))

	max = calMax(max, lh+rh+root.Val)

	return calMax(lh, rh) + root.Val
}

func calMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}