//2352. Equal Row and Column Pairs

func equalPairs(grid [][]int) int {
    n := len(grid)
    mapa := make(map[int][]int)
    result := 0

    for i, arr := range grid {
        hash := 0
        for _, v := range arr {
            hash += v
        } 
        _, ok := mapa[hash]
        if ok {
            mapa[hash] = append(mapa[hash], i)
            // fmt.Printf("add to exist hash %v\n", mapa[hash])
        } else {
            mapa[hash] = []int{i}
            // fmt.Printf("new hash %v\n", mapa[hash])
        }

    }

    for i := 0; i < n; i++ {
        curSum := 0
        curArr := make([]int,n)
        for j := 0; j < n; j ++ {
            curArr[j] = grid[j][i]
            curSum += curArr[j]
        }
        indexes, ok := mapa[curSum]
        if ok {
            for _, index := range indexes {
                if equals(curArr, grid[index]) {
                    result++
                }
            }
        }
    }

    return result
}

func equals(a1, a2 []int) bool {
    // fmt.Printf("checking %v and %v\n",a1, a2)
    for i := 0; i < len(a1); i++ {
        if a1[i] != a2[i] {
            return false
        }
    }
    return true
}

