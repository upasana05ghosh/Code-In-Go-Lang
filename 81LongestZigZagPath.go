//https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return max(findLongest(root.Left, true, 0), findLongest(root.Right, false, 0))
}

func findLongest(root *TreeNode, isLeft bool, count int) int {
	if root == nil {
		return count
	}

	//if left, go right with count +1 or go left and sart from 0
	if isLeft {
		return max(findLongest(root.Right, false, count+1), findLongest(root.Left, true, 0))
	}

	// for right, go left with count +1 or go right and start from 0
	return max(findLongest(root.Left, true, count+1), findLongest(root.Right, false, 0))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}