// https://leetcode.com/problems/missing-number/
// Given an array nums containing n distinct numbers in the range [0, n], return
// the only number in the range that is missing from the array.

package main

// time O(n), space O(1)
func missingNumber(nums []int) int {
	var sum1, sum2 int

	for i := 0; i < len(nums); i++ {
		sum1 += i
		sum2 += nums[i]
	}
	sum1 += len(nums)

	return sum1 - sum2
}

// time O(n), space O(1)
// using xor to find the odd one out
func missingNumber2(nums []int) int {
	result := 0

	for i, v := range nums {
		result ^= v ^ i
	}

	return result ^ len(nums)
}
