package main

// 5. Longest Palindromic Substring
// O(N^3)
func longestPalindromeSlow(s string) string {
	if len(s) == 1 {
		return s
	}
	var result string
	maxSubs := 0
	for i := len(s); i > 0; i-- {
		for j := 0; j <= len(s)-i; j++ {
			currentStr := s[j : j+i]
			if checkPalindromSlow(currentStr) {
				if len(currentStr) >= maxSubs {
					result = s[j : j+i]
					maxSubs = i
				}
			}
		}
	}
	return result
}

// O(N)
func checkPalindromSlow(s string) bool {
	left := 0
	right := len(s) - 1
	for right >= left {
		if s[left] != s[right] {
			return false
		}
		right--
		left++
	}
	return true
}
