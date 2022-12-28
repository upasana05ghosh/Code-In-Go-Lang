//https://leetcode.com/problems/binary-tree-paths/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var res = []string{}
	updatePath(root, &res, "")
	return res
}

func updatePath(root *TreeNode, res *[]string, sb string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*res = append(*res, sb+fmt.Sprintf("%d", root.Val))
		return
	}

	updatePath(root.Left, res, sb+fmt.Sprintf("%d", root.Val)+"->")
	updatePath(root.Right, res, sb+fmt.Sprintf("%d", root.Val)+"->")
}