package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := depth(node.Left)
		right := depth(node.Right)

		if left+right > maxDiameter {
			maxDiameter = left + right
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}

	depth(root)

	return maxDiameter
}
