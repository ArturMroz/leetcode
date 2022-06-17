// https://leetcode.com/problems/diameter-of-binary-tree/
// Given the root of a binary tree, return the length of the diameter of the tree.
// The diameter of a binary tree is the length of the longest path between any
// two nodes in a tree. This path may or may not pass through the root.

package main

// time O(n), space O(n)
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)
		maxDiameter = max(leftHeight+rightHeight, maxDiameter)

		return 1 + max(leftHeight, rightHeight)
	}

	_ = dfs(root)
	return maxDiameter
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * Definition for max func.
 * func max(a, b int) int {
 * 	if a > b {
 * 		return a
 * 	}
 * 	return b
 * }
 */
