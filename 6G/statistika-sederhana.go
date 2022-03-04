package main

import (
	"fmt"
)

func main() {
	var n, x int
	max := -100000
	min := 100000
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		if max < x {
			max = x
		}
		if min > x {
			min = x
		}
	}
	fmt.Print(max, min)
}
