package main

import "fmt"

func main() {
	var n, x, divide int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
	innerloop:
		for j := 1; j <= x-1; j++ {
			if divide > 3 {
				break innerloop
			} else if x%j == 0 {
				divide += 1
			}
		}
		if x == 1 {
			fmt.Println("YA")
		} else if divide <= 3 {
			fmt.Println("YA")
		} else {
			fmt.Println("BUKAN")
		}
		divide = 0
	}
}
