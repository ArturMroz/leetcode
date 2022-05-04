// https://leetcode.com/problems/single-number/
// Given a non-empty array of integers nums, every element appears twice except
// for one. Find that single one.
// You must implement a solution with a linear runtime complexity and use only
// constant extra space.

package main

// time O(n), space O(1)
func singleNumber(nums []int) int {
	result := 0

	// xoring two duplicates will evaluate to 0 (a XOR a == 0)
	// only the unique value will survive xor'ing together alllll the values
	for i := range nums {
		result ^= nums[i]
	}

	return result
}
