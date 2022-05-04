// https://leetcode.com/problems/maximum-depth-of-binary-tree/

// Given the root of a binary tree, return its maximum depth.
// A binary tree's maximum depth is the number of nodes along the longest path
// from the root node down to the farthest leaf node.

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// recursive solution
// time O(n), space O(n)
func maxDepthRec(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthRec(root.Left), maxDepthRec(root.Right))
}

// iterative solution
// time O(n), space O(n)
func maxDepthIter(root *TreeNode) int {
	type item struct {
		node  *TreeNode
		depth int
	}

	stack := []item{{root, 1}}
	maxLvl := 0

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur.node != nil {
			maxLvl = max(maxLvl, cur.depth)
			stack = append(stack, item{cur.node.Left, cur.depth + 1})
			stack = append(stack, item{cur.node.Right, cur.depth + 1})
		}
	}

	return maxLvl
}

func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}
