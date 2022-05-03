// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// You are given an array prices where prices[i] is the price of a given stock
// on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one stock
// and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction. If you
// cannot achieve any profit, return 0.

package main

// time O(n), space O(1)
func maxProfit(prices []int) int {
	l := 0
	r := 1
	max := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			cur := prices[r] - prices[l]
			if cur > max {
				max = cur
			}
		} else {
			l = r
		}
		r++
	}

	return max
}
