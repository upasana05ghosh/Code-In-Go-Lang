//https://leetcode.com/problems/binary-tree-preorder-traversal/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var res = []int{}        // res slice
	var stack = []TreeNode{} // stack slice
	var node TreeNode
	var n int

	for {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, *root)
			root = root.Left
		}

		if root == nil && len(stack) == 0 {
			return res
		}
		n = len(stack)
		node = stack[n-1]
		stack = stack[:n-1] //pop stack
		root = node.Right
	}

	return res
}