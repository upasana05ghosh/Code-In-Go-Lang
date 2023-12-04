//https://leetcode.com/problems/flatten-binary-tree-to-linked-list/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {
	var pre *TreeNode = nil

	helper(root, &pre)
}

func helper(root *TreeNode, pre **TreeNode) {
	if root == nil {
		return
	}

	helper(root.Right, pre)
	helper(root.Left, pre)

	root.Right = *pre
	root.Left = nil

	*pre = root
}