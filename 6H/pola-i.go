package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	for i := 1; i <= n; i++ {
		if i%k == 0 {
			fmt.Println("*")
		} else {
			fmt.Println(i)
		}
	}
}
