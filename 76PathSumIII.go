//https://leetcode.com/problems/path-sum-iii/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	valSet := make(map[int]int)
	valSet[0] = 1
	count := 0
	findCount(root, targetSum, valSet, &count, 0)
	return count
}

func findCount(root *TreeNode, targetSum int, valSet map[int]int, count *int, currSum int) {
	if root == nil {
		return
	}

	currSum += root.Val

	if v, ok := valSet[currSum-targetSum]; ok {
		*count += v
	}

	valSet[currSum] += 1

	findCount(root.Left, targetSum, valSet, count, currSum)
	findCount(root.Right, targetSum, valSet, count, currSum)

	valSet[currSum] -= 1
}