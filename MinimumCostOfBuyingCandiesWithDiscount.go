//2144. Minimum Cost of Buying Candies With Discount
func minimumCost(cost []int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(cost)))
    result := 0
    for i, v := range cost {
       if (i + 1) % 3 != 0 {
        result += v
       }
    }

    return result
}
