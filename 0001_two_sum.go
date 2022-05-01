// https://leetcode.com/problems/two-sum/
// Given an array of integers nums and an integer target, return indices of the
// two numbers such that they add up to target.
// Only one valid answer exists.

package main

// time: O(n), space: O(n)
func twoSum(nums []int, target int) []int {
	seen := map[int]int{}
	for i, v := range nums {
		if j, ok := seen[target-v]; ok {
			return []int{i, j}
		}
		seen[v] = i
	}

	return []int{}
}

// time: O(n^2), space: O(1)
func twoSum_bruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
