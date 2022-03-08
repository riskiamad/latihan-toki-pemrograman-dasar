package main

import (
	"fmt"
	"math"
)

var a, b, k, x float64

func f(a, x, b float64) float64 {
	return math.Abs(a*x + b)
}

func main() {

	fmt.Scan(&a, &b, &k, &x)
	for k > 0 {
		x = f(a, x, b)
		k--
	}
	fmt.Print(x)
}
