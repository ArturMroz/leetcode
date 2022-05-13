// https://leetcode.com/problems/number-of-1-bits/
// Write a function that takes an unsigned integer and returns the number of '1'
// bits it has (also known as the Hamming weight).

package main

// time O(1), space O(1)
func hammingWeight(num uint32) (result int) {
	for num != 0 {
		if num&1 != 0 {
			result++
		}
		num >>= 1
	}

	return result
}

// time O(1), space O(1)
func hammingWeight_cheeky(num uint32) (result int) {
	// this loop will only run as many times as many 1s are present in the uint
	// i.e. for 000001000010001 the loop will run only 3 times
	// because on each loop we're removing the rightmost '1' bit
	for num != 0 {
		num &= (num - 1)
		result++
	}

	return result
}
