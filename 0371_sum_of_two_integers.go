// https://leetcode.com/problems/sum-of-two-integers/
// Given two integers a and b, return the sum of the two integers without using
// the operators + and -.

package main

// time O(1), space O(1)
func getSum(a int, b int) int {
	for b != 0 {
		tmp := (a & b) << 1
		a ^= b
		b = tmp

		// or without using a temp variable:
		// a, b = a^b, (a&b)<<1
	}

	return a
}
