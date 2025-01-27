package main

import (
	"fmt"
)

func merge_sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left := merge_sort(arr[:len(arr)/2])
	right := merge_sort(arr[len(arr)/2:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	final_arr := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final_arr = append(final_arr, left[i])
			i++
		} else {

			final_arr = append(final_arr, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		final_arr = append(final_arr, left[i])
	}
	for ; j < len(right); j++ {
		final_arr = append(final_arr, right[j])
	}
	return final_arr
}

func main() {
	unsorted := []int{33, 22, 22, 33, 33, 10, 10, 6, 2, 22, 1, 5, 1, 5, 8, 3, 3, 4, 7, 9}
	// unsorted := []int{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sorted := merge_sort(unsorted)
	for i := range sorted {
		fmt.Printf(" %d ", sorted[i])
	}
}

type user struct {
	Email string
	HP    []byte
}
