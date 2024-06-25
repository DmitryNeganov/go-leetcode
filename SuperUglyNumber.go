//313. Super Ugly Number
func nthSuperUglyNumber(n int, primes []int) int {
    ugly := make([]int, n)
    ugly[0] = 1

    primesIndexes := make(map[int]int)
    for _, v := range primes {
        primesIndexes[v] = 0
    }

    for i := 1; i < n; i++ {
        next := math.MaxInt
        for _, v := range primes {
            cur := v * ugly[primesIndexes[v]]
            if cur < next {
                next = cur
            }
        }
        for _, v := range primes {
            cur := v * ugly[primesIndexes[v]]
            if next == cur {
                primesIndexes[v] = primesIndexes[v] + 1
            }
        }

        fmt.Println(next)
        ugly[i] = next
    }

    return ugly[n - 1]
}
