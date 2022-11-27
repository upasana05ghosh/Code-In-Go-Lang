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

	return checkSubtree(root.Left, root.Right)
}

func checkSubtree(left *TreeNode, right *TreeNode) bool {
	if left == nil {
		return right == nil
	}
	if right == nil {
		return left == nil
	}
	if left.Val != right.Val {
		return false
	}

	return checkSubtree(left.Right, right.Left) && checkSubtree(left.Left, right.Right)

}