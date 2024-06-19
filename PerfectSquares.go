//279. Perfect Squares
func numSquares(n int) int {
    sn := make([]int, n + 1)

    for i := 0; i < n + 1; i++ {
        sn[i] = -1
    }

    return dp(sn, n)
}

func dp(sn []int, n int) int {
    if n == 0 {
        sn[n] = 0
    }

    for i := 1; i * i <= n; i++ {
        if sn[n - i * i] == -1 {
            sn[n] = 1 + dp(sn, n - i * i)
        } else {
            sn[n] = min(1 + sn[n - i * i], sn[n])
        }
    }

    return sn[n]
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}
