//1791. Find Center of Star Graph
func findCenter(edges [][]int) int {
    mapa := make(map[int]int)

    for _, v := range edges {
        edge1 := v[0]
        edge2 := v[1]

        mapa[edge1] = mapa[edge1] + 1
        mapa[edge2] = mapa[edge2] + 1
    }

    n := len(mapa)

    for i := 1; i < n + 1; i++ {
        if mapa[i] == n - 1 {
            return i
        }
    }

    return -1
}
