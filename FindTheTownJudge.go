//997. Find the Town Judge

func findJudge(n int, trust [][]int) int {
    mapa := make(map[int]int)

    for i := 0; i < n + 1; i++ {
        mapa[i] = 0
    }

    for _, v := range trust {
        truster := v[0]
        judge := v[1]
        mapa[judge] = mapa[judge] + 1
        mapa[truster] = mapa[truster] - 1
    }

    for i := 1; i < n + 1; i++ {
        if mapa[i] == n - 1 {
            return i
        }
    }
    
    return - 1
}
