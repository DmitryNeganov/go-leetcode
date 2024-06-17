//718. Maximum Length of Repeated Subarray
func findLength(nums1 []int, nums2 []int) int {
    res := 0
    dp := make([][]int, len(nums1) + 1)
    for index := 0; index <= len(nums1); index++ {
        dp[index] = make([]int, len(nums2) + 1)
    }

    for i := len(nums1) - 1; i >= 0; i-- {
        for j := len(nums2) - 1; j >= 0; j--  {
            if nums1[i] == nums2[j] {
                dp[i][j] = dp[i + 1][j + 1] + 1
            }
            if dp[i][j] > res {
                res = dp[i][j]
            }
        }
    }

    return res
}
