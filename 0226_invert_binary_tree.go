// https://leetcode.com/problems/invert-binary-tree/
// Given the root of a binary tree, invert the tree, and return its root.

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive solution
// time O(n), space avg O(log n), space worst O(n)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// iterative solution
// time O(n), space avg O(log n), space worst O(n)
func invertTreeIterative(root *TreeNode) *TreeNode {
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			node.Left, node.Right = node.Right, node.Left
			stack = append(stack, node.Left, node.Right)
		}
	}

	return root
}
