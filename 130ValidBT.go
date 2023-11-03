//https://leetcode.com/problems/validate-binary-search-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var stack []*TreeNode
	var pre *TreeNode
	pre = nil
	for root != nil || len(stack) != 0 {

		//store val in left
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		//pop
		stack = stack[:len(stack)-1]

		if pre != nil && pre.Val >= root.Val {
			return false
		}

		pre = root
		root = root.Right
	}

	return true
}