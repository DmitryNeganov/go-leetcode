//523. Continuous Subarray Sum
func checkSubarraySum(nums []int, k int) bool {
    mapa := make(map[int]int)
    prev := 0
    for i, v := range nums{
        cur := (prev + v) % k
        _, ok := mapa[cur]
        if ok {
            return true
        }
        mapa[prev] = i
        prev = cur
    }
    return false
}
