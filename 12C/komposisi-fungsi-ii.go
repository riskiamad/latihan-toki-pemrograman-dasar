package main

import (
	"fmt"
	"math"
)

var a, b, k, x float64

func f(x float64) float64 {
	if k == 0 {
		return x
	}
	x = math.Abs(a*x + b)
	k--
	return f(x)
}

func main() {
	fmt.Scan(&a, &b, &k, &x)
	fmt.Print(f(x))
}
