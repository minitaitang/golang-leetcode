package interview

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBiNode(root *TreeNode) *TreeNode {
	newTree := &TreeNode{}
	inOrder(root, newTree)
	return newTree.Right
}

func inOrder(root *TreeNode, pos *TreeNode) *TreeNode {
	if root == nil {
		return pos
	}
	pos = inOrder(root.Left, pos)
	pos.Left = nil
	pos.Right = root
	pos = root
	pos = inOrder(root.Right, pos)

	return pos
}
