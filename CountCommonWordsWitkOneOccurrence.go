// 2085. Count Common Words With One Occurrence
func countWords(words1 []string, words2 []string) int {
    mapa := make(map[string]int)

    for _, str := range words1 {
        _, containsStr := mapa[str]
        if containsStr {
            mapa[str] = mapa[str] + 1
        } else {
            mapa[str] = 1
        }
    }

    counter := 0

    for _, str := range words2 {
        _, containsStr := mapa[str]
        if containsStr {
            if mapa[str] == 1 {
                mapa[str] = 0
                counter++
            } else if mapa[str] > 1 {
                delete(mapa, str)
            } else {
                counter--
                delete(mapa, str)
            }
        }
    }
    
    return counter
}
