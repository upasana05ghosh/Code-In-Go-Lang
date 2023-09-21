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
	var findLeaf1 []int
	var findLeaf2 []int
	findLeaf(root1, &findLeaf1)
	findLeaf(root2, &findLeaf2)

	fmt.Println(findLeaf1)
	fmt.Println(findLeaf2)
	return reflect.DeepEqual(findLeaf1, findLeaf2)
}

func findLeaf(root *TreeNode, leaf *[]int) {
	if root == nil {
		return
	}

	// for leave node
	if root.Left == nil && root.Right == nil {
		*leaf = append(*leaf, root.Val)
	}

	findLeaf(root.Left, leaf)
	findLeaf(root.Right, leaf)
}