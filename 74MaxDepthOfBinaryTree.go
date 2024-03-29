/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := maxDepth(root.Left)
	rh := maxDepth(root.Right)

	return max(lh, rh) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}