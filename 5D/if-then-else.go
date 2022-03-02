package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n > 0 {
		fmt.Print("positif")
	} else if n < 0 {
		fmt.Print("negatif")
	} else {
		fmt.Print("nol")
	}
}
