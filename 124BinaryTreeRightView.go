//https://leetcode.com/problems/binary-tree-right-side-view/description/
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

	var q []*TreeNode
	q = append(q, root)

	for len(q) != 0 {
		size := len(q)
		var n TreeNode
		for i := 0; i < size; i++ {
			n = *q[0]
			//remove first
			q = q[1:]

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		res = append(res, n.Val)
	}

	return res
}