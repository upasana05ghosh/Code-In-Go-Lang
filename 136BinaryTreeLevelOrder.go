//https://leetcode.com/problems/binary-tree-level-order-traversal/description
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var q []*TreeNode

	if root == nil {
		return res
	}

	q = append(q, root)

	for len(q) != 0 {
		size := len(q)
		var level []int
		for i := 0; i < size; i++ {
			n := q[0]
			//remove first
			q = q[1:]
			level = append(level, n.Val)

			if n.Left != nil {
				q = append(q, n.Left)
			}

			if n.Right != nil {
				q = append(q, n.Right)
			}
		}

		res = append(res, level)
	}

	return res

}