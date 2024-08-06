//1437. Check If All 1's Are at Least Length K Places Away
func kLengthApart(nums []int, k int) bool {
    var counter = math.MaxInt32

    for i := 0; i < len(nums); i++ {
        cur := nums[i]
        if cur == 1 {
            if counter < k {
                return false
            }
            counter = 0
        } else {
            counter++
        }
    }

    return true
}
