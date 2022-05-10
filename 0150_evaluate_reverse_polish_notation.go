// https://leetcode.com/problems/evaluate-reverse-polish-notation/
// Evaluate the value of an arithmetic expression in Reverse Polish Notation.
// Valid operators are +, -, *, and /. Each operand may be an integer or another expression.
//
// It is guaranteed that the given RPN expression is always valid:
// - the expression would always evaluate to a result
// - there won't be any division by zero operation.

package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, t := range tokens {
		switch t {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]

		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]

		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]

		case "/":
			// we are guaranteed by problem desc there won't be any division by zero
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]

		default:
			v, _ := strconv.Atoi(t)
			stack = append(stack, v)
		}
	}

	return stack[len(stack)-1]
}
