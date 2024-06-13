//209. Minimum Size Subarray Sum
func minSubArrayLen(target int, nums []int) int {
    left := 0
    sum := 0
    counter := math.MaxInt32 

    for right := 0; right < len(nums); right++ {
        sum += nums[right]
        for sum >= target {
            counter = min(counter, right - left + 1)
            sum -= nums[left]
            left++
        }
    }
    
    if counter == math.MaxInt32 {
        return 0
    } else {
        return counter
    }
}
