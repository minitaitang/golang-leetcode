package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSumIII(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := findPath437(root, sum)
	res += pathSumIII(root.Left, sum)
	res += pathSumIII(root.Right, sum)
	return res
}

func findPath437(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	res := 0
	if root.Val == sum {
		res++
	}
	res += findPath437(root.Left, sum-root.Val)
	res += findPath437(root.Right, sum-root.Val)

	return res
}
