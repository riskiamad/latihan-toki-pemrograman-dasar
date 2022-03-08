package main

import (
	"fmt"
	"math"
)

var n, x, y, d, min, max float64

func f(xi, xj, yi, yj float64) float64 {
	a := math.Abs(xj - xi)
	b := math.Abs(yj - yi)
	return math.Pow(a, d) + math.Pow(b, d)
}

func main() {
	fmt.Scan(&n, &d)
	var arr = make([][]float64, int(n))
	for i := 0; i < int(n); i++ {
		arr[i] = make([]float64, 2)
		fmt.Scan(&x, &y)
		arr[i][0] = x
		arr[i][1] = y
	}
	fmt.Print(arr)
	var data = []float64{}
	for i := 0; i < int(n); i++ {
		for j := 0; j < int(n); j++ {
			if i == j {
				continue
			}
			data = append(data, f(arr[i][0], arr[j][0], arr[i][1], arr[j][1]))
		}

	}

	for id, item := range data {
		if id == 0 {
			min = item
			max = item
		}
		if item > max {
			max = item
		}
		if item < min {
			min = item
		}
	}
	fmt.Print(min, max)
}
