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

// O(N2)
func longestPalindromeFaster(s string) string {

    n := len(s)

	iAns := 0
	jAns := 0

    dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
    
    for i := 0; i < n; i++ {
        dp[i][i] = true;
    }

    for i := 0; i < n - 1; i++ {
		dp[i][i+1] = (s[i] == s[i + 1])
        if dp[i][i+1] {
			iAns = i
			jAns = i + 1
		}
    }

	for dif := 2; dif < n; dif++ {
		for i := 0; i < n - dif ; i++ {
			j := i + dif
			dp[i][j] = (s[i] == s[j]) && (dp[i+1][j-1])
			if dp[i][j] {
				iAns = i
				jAns = j
			}
		}
	}

	return s[iAns : jAns + 1]
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
