// https://leetcode.com/problems/contains-duplicate/
// Given an integer array nums, return true if any value appears at least twice
// in the array, and return false if every element is distinct.

package main

func containsDuplicate(nums []int) bool {
	unique := map[int]struct{}{}
	dummy := struct{}{}

	for _, v := range nums {
		unique[v] = dummy
	}

	return len(unique) != len(nums)
}
