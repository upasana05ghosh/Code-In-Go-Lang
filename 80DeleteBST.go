//https://leetcode.com/problems/delete-node-in-a-bst/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, k int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > k {
		root.Left = deleteNode(root.Left, k)
	} else if root.Val < k {
		root.Right = deleteNode(root.Right, k)
	} else { //node to delete

		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		minNode := findMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, root.Val)
	}
	return root
}

func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}