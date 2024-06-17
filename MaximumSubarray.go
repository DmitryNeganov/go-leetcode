//53. Maximum Subarray
func maxSubArray(nums []int) int {
    maxSum := math.MinInt
    sum := 0

    for _, v := range nums {

        curSum := sum + v
        if curSum < 0 {
            sum = 0
        } else {
            sum = curSum
        }

        if curSum > maxSum {
            maxSum = curSum
        }
    }

    return maxSum
}
