package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	stack = append(stack, p, q)
	for len(stack) != 0 {
		left, right := stack[len(stack)-1], stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		stack = append(stack, left.Left)
		stack = append(stack, right.Left)
		stack = append(stack, left.Right)
		stack = append(stack, right.Right)
	}
	return true
}
