//   https://leetcode.com/problems/reverse-linked-list/
//   Given the head of a singly linked list, reverse the list.

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// iterative solution
// time O(n), space O(1)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next

		// oneliner version of the swap party above
		// head.Next, prev, head = prev, head, head.Next
	}

	return prev
}

// recursive solution
// time O(n), space O(n)
func reverseListRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListRec(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
