//263. Ugly Number
func isUgly(n int) bool {
    if n < 1 {
        return false
    }

    for n > 1 {
        if n % 5 == 0 {
            n = n / 5
        } else if n % 3 == 0 {
            n = n / 3
        } else if n % 2 == 0 {
            n = n / 2
        } else {
            return false
        }
    }

    return true
}
