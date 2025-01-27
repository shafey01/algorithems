package main

import (
	"fmt"
	"slices"
)

func main() {
	dp := make([]int, 8)
	p := []int{8, 3, 5, 2, 4, 9, 7, 11}
	n := len(p)
	for i := n - 1; i >= 0; i-- {

		ch := []int{1}
		for j := i + 1; j < n; j++ {
			if p[j] > p[i] {
				ch = append(ch, dp[j]+1)
			}
		}
		dp[i] = slices.Max(ch)
	}
	fmt.Println(slices.Max(dp))
}
