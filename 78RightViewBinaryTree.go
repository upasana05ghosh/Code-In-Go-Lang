//https://leetcode.com/problems/binary-tree-right-side-view/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	var list []*TreeNode
	list = append(list, root)

	for len(list) != 0 {
		size := len(list)
		var n *TreeNode
		for i := 0; i < size; i++ {
			n = list[0]

			//remove this elemenet from slice
			list = list[1:]

			if n.Left != nil {
				list = append(list, n.Left)
			}

			if n.Right != nil {
				list = append(list, n.Right)
			}
		}
		res = append(res, n.Val)
	}

	return res
}