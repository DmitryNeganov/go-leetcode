package main

import "fmt"

func minMaxGame(nums []int) int {
	n := len(nums)
	if n == 1 {
		fmt.Println(nums[0])
		return nums[0]
	}
	newNums := make([]int, n/2, n/2)

	for i := 0; i < len(newNums); i++ {
		if i%2 == 0 {
			newNums[i] = min(nums[2*i], nums[2*i+1])
		} else {
			newNums[i] = max(nums[2*i], nums[2*i+1])
		}
	}

	return minMaxGame(newNums)
}
