//49. Group Anagrams

type Key [26]int

func getKey(str string) (key Key) {
    for i := range str {
        key[str[i] - 'a']++
    }
    return
}

func groupAnagrams(strs []string) [][]string {
    mapa := make(map[Key][]string)

    for _, v := range strs {
        key := getKey(v)
        mapa[key] = append(mapa[key], v)
    }

    result := make([][]string, 0, len(mapa))

    for _, v := range mapa {
        result = append(result, v)
    }

    return result
}
