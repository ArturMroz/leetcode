// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Given the head of a linked list, remove the nth node from the end of the list
// and return its head. Do this in one pass.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

// time O(n), space O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head} // attaching dummy at start so it's easier to remove head
	left := dummy
	right := head

	for n > 0 && right != nil {
		right = right.Next
		n--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next

	return dummy.Next
}

// time O(n), space O(1)
func removeNthFromEnd_alt(head *ListNode, n int) *ListNode {
	right := head
	left := head
	count := 0

	for right != nil {
		if count > n {
			left = left.Next
		}
		count++
		right = right.Next
	}

	// only need to remove the head
	if n == count {
		return head.Next
	}

	left.Next = left.Next.Next

	return head
}
