//https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var q []*TreeNode
	q = append(q, root)
	max := math.MinInt32
	level := 0
	maxLevel := 0

	for len(q) != 0 {
		size := len(q)
		sum := 0
		level++
		for i := 0; i < size; i++ {
			first := q[0]
			sum += first.Val

			//remove first
			q = q[1:]

			if first.Left != nil {
				q = append(q, first.Left)
			}

			if first.Right != nil {
				q = append(q, first.Right)
			}
		}

		if sum > max {
			max = sum
			maxLevel = level
		}
	}

	return maxLevel
}