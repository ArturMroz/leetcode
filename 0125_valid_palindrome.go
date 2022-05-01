// https://leetcode.com/problems/valid-palindrome/
// Given a string s, return true if it is a palindrome, or false otherwise.
// s consists only of printable ASCII characters.

package main

// time: O(n), space: O(n)
func isPalindrome(s string) bool {
	cleaned := make([]byte, 0, len(s))

	for _, v := range []byte(s) { // remove all nonalphanumeric chars
		if 'A' <= v && v <= 'Z' {
			cleaned = append(cleaned, v+32) // if uppercase, convert to lowercase
		} else if ('a' <= v && v <= 'z') || ('0' <= v && v <= '9') {
			cleaned = append(cleaned, v)
		}
	}

	for i, j := 0, len(cleaned)-1; i < j; i, j = i+1, j-1 {
		if cleaned[i] != cleaned[j] {
			return false
		}
	}

	return true
}

// time: O(n), space: O(1)
func isPalindrome_withoutHelperArray(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if !isAlphanum(s[i]) {
			i++
		} else if !isAlphanum(s[j]) {
			j--
		} else {
			if toLower(s[i]) != toLower(s[j]) {
				return false
			}
			i++
			j--
		}
	}

	return true
}

func isAlphanum(v byte) bool {
	return ('a' <= v && v <= 'z') || ('A' <= v && v <= 'Z') || ('0' <= v && v <= '9')
}

func toLower(v byte) byte {
	if 'A' <= v && v <= 'Z' {
		return v + 32
	}
	return v
}
