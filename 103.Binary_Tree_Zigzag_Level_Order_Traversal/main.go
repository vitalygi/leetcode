package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	direction := true
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			currentNode := queue[i]

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}

			if direction {
				currentLevel = append(currentLevel, currentNode.Val)
			} else {
				currentLevel = append([]int{currentNode.Val}, currentLevel...)
			}
		}

		queue = queue[levelSize:]
		direction = !direction
		result = append(result, currentLevel)
	}

	return result
}
