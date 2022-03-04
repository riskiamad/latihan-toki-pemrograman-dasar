package main

import "fmt"

func main() {
	var n, x int
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		var bil int
		fmt.Scan(&bil)
		x += bil
	}
	fmt.Printf("%d", x)
}
