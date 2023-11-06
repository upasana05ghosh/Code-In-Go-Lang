//https://leetcode.com/problems/binary-tree-inorder-traversal/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int

	findInorder(root, &res)

	return res
}

func findInorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	findInorder(root.Left, res)
	*res = append(*res, root.Val)
	findInorder(root.Right, res)
}
