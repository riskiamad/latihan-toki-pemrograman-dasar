package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := n; i >= 1; i-- {
		for j := 1; j <= n; j++ {
			if j < i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}
}
