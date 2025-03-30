package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if (targetSum-root.Val) == 0 && root.Right == nil && root.Left == nil {
		return true
	}
	leftSum := hasPathSum(root.Left, targetSum-root.Val)
	rightSum := hasPathSum(root.Right, targetSum-root.Val)
	return leftSum == true || rightSum == true
}
