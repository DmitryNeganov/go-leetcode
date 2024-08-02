//1941. Check if All Characters Have Equal Number of Occurrences
func areOccurrencesEqual(s string) bool {
    mapa := make(map[rune]int)
    runa := []rune(s)

    for _, r := range runa {
        _, ok := mapa[r]
        if ok {
            mapa[r] = mapa[r] + 1
        } else {
            mapa[r] = 1
        }
    } 

    appearTimes := mapa[runa[0]]

    for _, v := range mapa {
        if v != appearTimes {
            return false
        }
    }

    return true
}
