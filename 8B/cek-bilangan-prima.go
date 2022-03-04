package main

import "fmt"

func main() {
	var n, x, divide int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		for j := 1; j <= 1000; j++ {
			if x == j {
				continue
			} else if x%j == 0 {
				divide += 1
			}
		}
		if x == 1 {
			fmt.Println("BUKAN")
		} else if divide == 1 {
			fmt.Println("YA")
		} else {
			fmt.Println("BUKAN")
		}
		divide = 0
	}
}
