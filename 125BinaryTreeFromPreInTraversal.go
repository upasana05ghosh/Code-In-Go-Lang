//https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return makeTree(0, 0, len(inorder), preorder, inorder)
}

func makeTree(preStart int, inStart int, inEnd int, preorder []int, inorder []int) *TreeNode {
	if preStart > len(preorder)-1 || inStart > inEnd {
		return nil
	}
	root := &TreeNode{Val: preorder[preStart]}

	//find root val in inorder
	var inIndex int
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == root.Val {
			inIndex = i
			break
		}
	}

	root.Left = makeTree(preStart+1, inStart, inIndex-1, preorder, inorder)
	root.Right = makeTree(preStart+inIndex-inStart+1, inIndex+1, inEnd, preorder, inorder)

	return root
}