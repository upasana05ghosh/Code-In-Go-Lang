//https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		//pop
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return root.Val
		}

		root = root.Right
	}
}