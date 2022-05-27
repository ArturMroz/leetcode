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
			for j := l; j < v; j++ {
				delete(seen, s[j])
			}
			l = v + 1
		}

		seen[s[i]] = i
		if len(seen) > max {
			max = len(seen)
		}
	}

	return max
}

// time O(n), space O(n)
func lengthOfLongestSubstring_flat(s string) int {
	max := 0
	seen := map[byte]struct{}{}
	l, r := 0, 0

	for l < len(s) && r < len(s) {
		if _, ok := seen[s[r]]; ok {
			delete(seen, s[l])
			l++
		} else {
			seen[s[r]] = struct{}{}
			r++
			if r-l > max {
				max = r - l
			}
		}
	}

	return max
}
