package main

import "fmt"

var n int

func biner(n int) string {
	if n == 1 {
		return "1"
	} else if n%2 == 1 {
		return biner(n/2) + "1"
	} else {
		return biner(n/2) + "0"
	}
}

func main() {
	fmt.Scan(&n)
	fmt.Print(biner(n))
}
