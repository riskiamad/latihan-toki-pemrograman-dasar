package main

import (
	"fmt"
)

func main() {
	var n float64
	fmt.Scan(&n)
	for n >= 1 {
		if n == 1 {
			fmt.Println("ya")
		} else if n < 2 && n != 1 {
			fmt.Println("bukan")
		}
		n = n / 2
	}
}
