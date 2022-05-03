// https://leetcode.com/problems/valid-parentheses/
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
//
// An input string is valid if:
// 1. Open brackets must be closed by the same type of brackets.
// 2. Open brackets must be closed in the correct order.

package main

// time O(n), space O(n)
func parens_isValid(s string) bool {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		switch ch := s[i]; ch {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			if len(stack) == 0 || stack[len(stack)-1] != ch {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// time O(n), space O(n)
func parens_isValid_withMap(s string) bool {
	stack := []byte{}
	parens := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := 0; i < len(s); i++ {
		switch ch := s[i]; ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		default:
			if len(stack) == 0 || stack[len(stack)-1] != parens[ch] {
				return false
			}
			stack = stack[:len(stack)-1]

		}
	}

	return len(stack) == 0
}
