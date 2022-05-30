// https://leetcode.com/problems/add-two-numbers/
// You are given two non-empty linked lists representing two non-negative
// integers. The digits are stored in reverse order, and each of their nodes
// contains a single digit. Add the two numbers and return the sum as a linked
// list.
// You may assume the two numbers do not contain any leading zero, except the
// number 0 itself.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

// time O(n), space O(n), where n is lenght of list with more nodes
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0
	v1, v2 := 0, 0

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}

		n := v1 + v2 + carry
		n, carry = n%10, n/10

		cur.Next = &ListNode{Val: n}
		cur = cur.Next
	}

	return dummy.Next
}
