// Package word provides utilities for word games.
package word

// IsPalindrome reports whether s reads the same forwards and backwards.
// (Our first attempt.)
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}