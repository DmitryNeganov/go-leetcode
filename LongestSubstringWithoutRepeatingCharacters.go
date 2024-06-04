//3. Longest Substring Without Repeating Characters
func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }

    r := []rune(s)
    mapa := make(map[rune]int)
    longest := 1
    currentLength := 1
    mapa[r[0]] = 0
    left := 0

    for i := 1; i < len(s); i++ {
        cur := r[i]
        position, ok := mapa[cur]
        if ok && position >= left {
            left = position
            currentLength = i - left
            mapa[cur] = i
        } else {
            currentLength++
            mapa[cur] = i
        }
        if longest < currentLength {
            longest = currentLength
        }
    }

    return longest
}
