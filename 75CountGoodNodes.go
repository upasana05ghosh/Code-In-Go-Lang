//https://leetcode.com/problems/count-good-nodes-in-binary-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	count := 0
	if root == nil {
		return count
	}
	getCount(root, root.Val, &count)
	return count
}

func getCount(root *TreeNode, max int, count *int) {
	if root == nil {
		return
	}
	if root.Val >= max {
		*count = *count + 1
		max = root.Val
	}

	getCount(root.Left, max, count)
	getCount(root.Right, max, count)
}