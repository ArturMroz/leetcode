// https://leetcode.com/problems/counting-bits/
// Given an integer n, return an array ans of length n + 1 such that for each i
// (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

package main

// time O(n), space O(n)
func countBits(n int) []int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = dp[i>>1] + (i & 1)
	}

	return dp
}

// time O(n), space O(n)
func countBits2(n int) []int {
	dp := make([]int, n+1)
	offset := 1

	for i := 1; i <= n; i++ {
		if offset*2 == i {
			offset = i // set offset to current most significant bit
		}
		dp[i] = 1 + dp[i-offset]
	}

	return dp
}
