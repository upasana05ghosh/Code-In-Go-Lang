/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var max int

func diameterOfBinaryTree(root *TreeNode) int {
	max = 0
	height(root)
	return max
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := height(root.Left)
	rh := height(root.Right)

	max = findMax(max, lh+rh)

	return findMax(lh, rh) + 1
}

func findMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}