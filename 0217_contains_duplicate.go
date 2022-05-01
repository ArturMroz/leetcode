// https://leetcode.com/problems/contains-duplicate/
// Given an integer array nums, return true if any value appears at least twice
// in the array, and return false if every element is distinct.

package main

// time: O(n), space: O(n)
func containsDuplicate(nums []int) bool {
	unique := map[int]struct{}{}

	for _, v := range nums {
		unique[v] = struct{}{}
	}

	return len(unique) != len(nums)
}
