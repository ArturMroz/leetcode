// https://leetcode.com/problems/reverse-bits/
// Reverse bits of a given 32 bits unsigned integer.

package main

// time O(1), space O(1)
func reverseBits(num uint32) (result uint32) {
	for i := 0; i < 31; i++ {
		result |= num & 1
		result <<= 1
		num >>= 1
	}

	result |= num & 1
	return result
}

// time O(1), space O(1)
func reverseBits2(num uint32) (result uint32) {
	for i := 0; i < 32; i++ {
		bit := (num >> i) & 1
		result |= bit << (31 - i)
	}

	return result
}
