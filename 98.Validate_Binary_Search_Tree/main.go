package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	type stackFrame struct {
		node     *TreeNode
		maxLimit *int
		minLimit *int
	}

	stack := []stackFrame{{node: root, maxLimit: nil, minLimit: nil}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.node == nil {
			continue
		}

		if (current.maxLimit != nil && current.node.Val >= *current.maxLimit) ||
			(current.minLimit != nil && current.node.Val <= *current.minLimit) {
			return false
		}

		currentVal := current.node.Val

		stack = append(stack, stackFrame{
			node:     current.node.Right,
			maxLimit: current.maxLimit,
			minLimit: &currentVal,
		})

		stack = append(stack, stackFrame{
			node:     current.node.Left,
			maxLimit: &currentVal,
			minLimit: current.minLimit,
		})
	}

	return true
}
