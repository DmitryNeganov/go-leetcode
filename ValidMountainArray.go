//941. Valid Mountain Array

func validMountainArray(arr []int) bool {
    if len(arr) < 3 {
        return false
    }

    increasing := true
    decreasing := false

    if arr[0] >= arr[1] {
        return false
    }

    for i := 2; i < len(arr); i++ {
        cur := arr[i]
        prev := arr[i-1]
        if prev == cur {
            return false
        }
        if increasing && prev > cur {
            increasing = false
            decreasing = true
        } 
        if decreasing && prev < cur {
            return false
        }
    }

    return decreasing
}
