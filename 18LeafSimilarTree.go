//https://leetcode.com/problems/leaf-similar-trees/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var s1, s2 string

	udpateListWithLeafNode(root1, &s1)
	udpateListWithLeafNode(root2, &s2)
	return s1 == s2
}

func udpateListWithLeafNode(root *TreeNode, s *string) {
	if root == nil {
		return
	}

	udpateListWithLeafNode(root.Left, s)

	if root.Left == nil && root.Right == nil {
		*s += fmt.Sprintf("%d ", root.Val)
	}

	udpateListWithLeafNode(root.Right, s)
}