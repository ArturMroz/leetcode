// https://leetcode.com/problems/binary-search/
// Given an array of integers nums which is sorted in ascending order, and an
// integer target, write a function to search target in nums. If target exists,
// then return its index. Otherwise, return -1.

package main

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		i := (start + end) / 2
		if nums[i] == target {
			return i
		}

		if nums[i] > target {
			end = i - 1
		} else {
			start = i + 1
		}
	}

	return -1
}
