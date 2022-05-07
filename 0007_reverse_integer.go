// https://leetcode.com/problems/reverse-integer/
// Given a signed 32-bit integer x, return x with its digits reversed. If
// reversing x causes the value to go outside the signed 32-bit integer range
// [-231, 231 - 1], then return 0.
//
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

package main

const (
	maxInt32 = 1<<31 - 1
	minInt32 = -1 << 31
)

func reverse(x int) (result int) {
	for x != 0 {
		digit := x % 10

		// overflow check (we can only use 32-bit integers in this problem)
		if (result > maxInt32/10) || (result == maxInt32/10 && digit > maxInt32%10) {
			return 0
		}
		if (result < minInt32/10) || (result == minInt32/10 && digit < minInt32%10) {
			return 0
		}

		result = result*10 + digit
		x /= 10
	}

	return result
}
