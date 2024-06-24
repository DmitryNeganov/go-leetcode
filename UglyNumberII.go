//264. Ugly Number II
func nthUglyNumber(n int) int {
    ugly := make([]int, n)
    ugly[0] = 1
    i2 := 0
    i3 := 0
    i5 := 0

    for i := 1; i < n; i++ {

        next := min(2 * ugly[i2], min(3 * ugly[i3], 5 * ugly[i5]))

        if next == 2 * ugly[i2] {
            i2++
        }
        if next == 3 * ugly[i3] {
            i3++
        }
        if next == 5 * ugly[i5] {
            i5++
        }

        ugly[i] = next
    }

    return ugly[n - 1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
