package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)
	sisa := num1 % num2
	masing := num1 / num2
	fmt.Println("masing-masing ", masing)
	fmt.Println("bersisa ", sisa)
}
