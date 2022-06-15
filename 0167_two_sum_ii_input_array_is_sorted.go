// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// Given a 1-indexed array of integers numbers that is already sorted in
// non-decreasing order, find two numbers such that they add up to a specific
// target number.
// There is exactly one solution. You may not use the same element twice.
// Your solution must use only constant extra space.

package main

// time O(n), space O(1)
func twoSum2_twoPointers(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1

	// problem description guarantees that a solution exsists, loop has to end at some point
	for numbers[start]+numbers[end] != target {
		if numbers[start]+numbers[end] > target {
			end--
		} else {
			start++
		}
	}

	return []int{start + 1, end + 1}
}

// time O(n log n), space O(1)
func twoSum2_binarySearch(numbers []int, target int) []int {
	end := len(numbers) - 1

	for i := 0; i < len(numbers); i++ {
		t := target - numbers[i]
		start := i + 1
		for start <= end {
			mid := (end + start) / 2
			cur := numbers[mid]
			if cur == t {
				return []int{i + 1, mid + 1}
			}
			if cur > t {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	return []int{0, 0}
}

// time O(n^2), space O(1)
func twoSum2_bruteForce(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			} else if numbers[i]+numbers[j] > target {
				break
			}
		}
	}

	return []int{0, 0}
}
