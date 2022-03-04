package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j <= i {
				fmt.Print(x)
				x += 1
				if x > 9 {
					x = 0
				}
			}
		}
		fmt.Print("\n")
	}
}
