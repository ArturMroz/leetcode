// https://leetcode.com/problems/climbing-stairs/
// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps.
// In how many distinct ways can you climb to the top?

package main

// time O(n), space O(1)
func climbStairs(n int) int {
	one := 1
	two := 1

	for i := 0; i < n-1; i++ {
		one, two = one+two, one
	}

	return one
}

// time O(n), space O(n)
func climbStairs_array(n int) int {
	dp := []int{1, 1}
	for i := 2; i <= n; i++ {
		dp = append(dp, dp[i-2]+dp[i-1])
	}
	return dp[n]
}
