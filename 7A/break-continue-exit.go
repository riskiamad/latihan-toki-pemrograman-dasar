package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%10 == 0 {
			continue
		} else if i == 42 {
			fmt.Println("ERROR")
			break
		}
		fmt.Println(i)
	}
}
