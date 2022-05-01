// https://leetcode.com/problems/valid-anagram/
// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// s and t consist of lowercase English letters.

package main

import "sort"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// s & and consist only of lowercase English letters (26 total)
	arr := make([]int, 26)
	for i := range s {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}

	for _, v := range arr {
		if v != 0 {
			return false
		}
	}

	return true
}

func isAnagram_sort(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sb := []byte(s)
	sort.Slice(sb, func(i, j int) bool { return sb[i] < sb[j] })

	tb := []byte(t)
	sort.Slice(tb, func(i, j int) bool { return tb[i] < tb[j] })

	return string(sb) == string(tb)
}
