// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// Given a string s, find the length of the longest substring without repeating characters.

package main

// time O(n), space O(n)
func lengthOfLongestSubstring(s string) int {
	seen := map[byte]int{}
	max := 0
	l := 0

	for i := range s {
		if v, ok := seen[s[i]]; ok {
			for j := l; j <= v; j++ {
				delete(seen, s[j])
				l++
			}
		}

		seen[s[i]] = i
		if len(seen) > max {
			max = len(seen)
		}
	}

	return max
}
