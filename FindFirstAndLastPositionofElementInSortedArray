//34. Find First and Last Position of Element in Sorted Array

func searchRange(nums []int, target int) []int {
    result := make([]int, 0, 2)
    result = append(result, binSearch(nums, target, true))
    result = append(result, binSearch(nums, target, false))
    return result
}

func binSearch(nums []int, target int, isLeftBound bool) int {
    left := 0
    right := len(nums) - 1
    res := -1
    for left <= right {
        mid := (left + right) / 2
        cur := nums[mid]
        if cur == target {
            if isLeftBound {
                right = mid - 1
                res = mid
            } else {
                left = mid + 1
                res = mid
            }
        } else if cur > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return res
}
