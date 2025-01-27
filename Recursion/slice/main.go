package main

import (
	"fmt"
	"slices"
)

func reverse_slice(x []int) []int {
	if len(x) == 1 {
		return x
	}

	t := x[0]
	x = x[1:]

	reverse_slice(x)
	x = append(x, t)
	return x
	// x := append(x, x[0])
}

// func reverse_slice_in_slice(x [] int ) [] int {
//
// }
func reverse_looop(x []int) {
	for i, j := 0, len(x)-1; i < 1; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
}

func main() {
	s := []int{1, 2, 3, 4}
	reverse_slice(s)
	slices.Reverse(s)
	fmt.Println(s)
	// fmt.Println(reverse_slice_in_slice(r))
}
