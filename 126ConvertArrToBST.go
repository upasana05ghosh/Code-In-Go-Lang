//https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(a []int) *TreeNode {
	if len(a) == 0 {
		return nil
	}

	return bs(a, 0, len(a)-1)
}

func bs(a []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2

	root := &TreeNode{
		Val:   a[mid],
		Left:  bs(a, start, mid-1),
		Right: bs(a, mid+1, end),
	}
	return root

}