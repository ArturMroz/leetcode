// https://leetcode.com/problems/find-the-duplicate-number/
// Given an array of integers nums containing n + 1 integers where each integer
// is in the range [1, n] inclusive.
// There is only one repeated number in nums, return this repeated number.
// You must solve the problem without modifying the array nums and uses only
// constant extra space.

package main

// time O(n), space O(1)
func findDuplicate(nums []int) int {
	// treating nums as linked list and using Floyd's algo to find cycle in linked list
	slow := 0
	fast := 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			return slow
		}
	}
}

// time O(n log n), space O(1)
func findDuplicate_binarySearch(nums []int) int {
	low := 1
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		cnt := 0
		for _, a := range nums {
			if a <= mid {
				cnt++
			}
		}

		if cnt <= mid {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

// time O(n), space O(n)
func findDuplicate_map(nums []int) int {
	seen := map[int]bool{}
	for _, v := range nums {
		if dup := seen[v]; dup {
			return v
		}
		seen[v] = true
	}

	return -1
}
