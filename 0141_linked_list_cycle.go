// https://leetcode.com/problems/linked-list-cycle/
// Determine if the linked list has a cycle in it.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

// time O(n), space O(1)
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// tortoise and hare algorithm: if tortoise caught up with hare, we've got a cycle
		if slow == fast {
			return true
		}
	}

	return false
}

// time O(n), space O(n)
func hasCycle_usingMap(head *ListNode) bool {
	seen := make(map[*ListNode]struct{})

	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := seen[cur]; ok {
			return true
		}
		seen[cur] = struct{}{}
	}

	return false
}
