package main

import (
	"fmt"
	"time"
)

func main() {
	res := map[int]int{
		1: 1,
		2: 1,
	}
	x := 100

	now := time.Now()
	// a := fibo(x)
	fmt.Printf("Fibo takes: %v time \n", time.Now().Sub(now))

	now2 := time.Now()
	b := fibo_efficient(x, res)
	fmt.Printf("Fibo efficent takes: %v time \n", time.Now().Sub(now2))

	// fmt.Println(a)
	fmt.Println(b)
	fmt.Println(fo(9000000))
}

func fibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fibo(n-1) + fibo(n-2)
}

func fibo_efficient(n int, res map[int]int) int {
	ans, ok := res[n]
	if ok {
		return ans
	} else {
		ans = fibo_efficient(n-1, res) + fibo_efficient(n-2, res)
		res[n] = ans
		return ans
	}
}

func fo(x int) int {
	return x * x
}
