//O(N3)
func countSubstringsSlow(s string) int {
    counter := 0
    for i := 0; i < len(s); i++ {
        for j := i + 1; j < len(s) + 1; j++ {
            if isPalindrome(s[i:j]) {
                counter++
            }
        }
    }
    return counter
}

//O(N)
func isPalindrome(s string) bool {
    left := 0
    right := len(s) - 1

    for left <= right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}
