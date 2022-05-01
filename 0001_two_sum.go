// https://leetcode.com/problems/two-sum/
// Given an array of integers nums and an integer target, return indices of the
// two numbers such that they add up to target.

package main

func twoSum(nums []int, target int) []int {
	hashmap := map[int]int{}
	for i, v := range nums {
		if j, ok := hashmap[target-v]; ok {
			return []int{i, j}
		}
		hashmap[v] = i
	}

	return []int{}
}

func twoSum_bruteForce(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
