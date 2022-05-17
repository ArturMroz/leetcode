// https://leetcode.com/problems/balanced-binary-tree/submissions/
// Given a binary tree, determine if it is height-balanced.
// For this problem, a height-balanced binary tree is defined as a binary tree
// in which the left and right subtrees of every node differ in height by no
// more than 1.

package main

// time O(n), space avg O(log n), space worst O(n)
func isBalanced(root *TreeNode) bool {
	_, balanced := isBalancedHelper(root)
	return balanced
}

func isBalancedHelper(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftH, leftOk := isBalancedHelper(node.Left)
	rightH, rightOk := isBalancedHelper(node.Right)
	diff := leftH - rightH

	balanced := leftOk && rightOk && (-2 < diff && diff < 2)

	maxH := leftH
	if rightH > leftH {
		maxH = rightH
	}

	return 1 + maxH, balanced
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
