// https://leetcode.com/problems/reorder-list/
// You are given the head of a singly linked-list:
// L0 → L1 → … → Ln - 1 → Ln
// Reorder the list to be on the following form:
// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

// time O(n), space O(1)
func reorderList(head *ListNode) {
	// find the middle: fast pointer goes 2x faster than the slow one;
	// when the fast one is finished, the slow one will point at the
	// middle of the list
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	// the end of the first part of the list should now point at nil
	slow.Next = nil

	// reverse the 2nd part of the list
	var prev *ListNode = nil
	for second != nil {
		tmp := second.Next
		second.Next = prev
		prev = second
		second = tmp
	}

	// zip two lists together
	list1 := head
	list2 := prev

	for list2 != nil {
		tmp1 := list1.Next
		tmp2 := list2.Next

		list2.Next = list1.Next
		list1.Next = list2

		list1 = tmp1
		list2 = tmp2
	}
}

// time O(n), space O(n)
func reorderList_array(head *ListNode) {
	nodes := []*ListNode{}

	cur := head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	for i := 0; i < len(nodes)/2; i++ {
		j := len(nodes) - 1 - i
		nodes[j].Next = nodes[i].Next
		nodes[i].Next = nodes[j]
	}

	nodes[len(nodes)/2].Next = nil
}
