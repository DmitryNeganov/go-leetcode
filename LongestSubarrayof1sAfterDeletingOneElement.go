//1493. Longest Subarray of 1's After Deleting One Element
func longestSubarray(nums []int) int {
    left := 0
    i := 0
    zeroPos := -1
    longest := 0

    for i < len(nums) {
        cur := nums[i]
        if cur == 0 {
            if zeroPos < left {
                zeroPos = i
            } else {
                left = zeroPos + 1
                zeroPos = i
            }
        }
        longest = max(longest, i - left)
        i++
    }

    return longest
}
