package main

import "fmt"

var n, resopitory, res, result int

func faktorisasi(n int) int {
	if n == 1 {
		return n
	} else if n%2 == 0 {
		return n / 2 * faktorisasi(n-1)

	} else {
		return n * faktorisasi(n-1)
	}

}

func main() {
	fmt.Scan(&n)
	fmt.Print(faktorisasi(n))
}
