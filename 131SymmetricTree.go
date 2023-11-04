//https://leetcode.com/problems/symmetric-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return findSymmetric(root.Left, root.Right)
}

func findSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil {
		return right == nil
	}

	if right == nil {
		return left == nil
	}

	if left.Val != right.Val {
		return false
	}

	return findSymmetric(left.Left, right.Right) && findSymmetric(left.Right, right.Left)
}