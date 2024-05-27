//2073. Time Needed to Buy Tickets
func timeRequiredToBuy(tickets []int, k int) int {
    counter := 0

    i := 0
    for tickets[k] > 0 {
        current := tickets[i]
        if current > 0 {
            counter++
            tickets[i] = current - 1
        }
        if i == len(tickets) - 1 {
            i = 0
        } else {
            i++
        }
    }

    return counter
}
