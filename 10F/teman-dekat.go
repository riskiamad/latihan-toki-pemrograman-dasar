package main

import (
	"fmt"
	"math"
)

var n, x, y, min, max int64
var d float64

func f(xi, xj, yi, yj int64) int64 {
	a := math.Abs(float64(xj - xi))
	b := math.Abs(float64(yj - yi))
	return int64(math.Pow(a, d) + math.Pow(b, d))
}

func main() {
	fmt.Scan(&n, &d)
	var arr = make([][]int64, int(n))
	for i := 0; i < int(n); i++ {
		arr[i] = make([]int64, 2)
		fmt.Scan(&x, &y)
		arr[i][0] = x
		arr[i][1] = y
	}
	i := 0
	j := 1
	var data = []int64{}
	for i < len(arr)-1 {
		data = append(data, f(arr[i][0], arr[j][0], arr[i][1], arr[j][1]))

		if j < len(arr)-1 {
			j += 1
		} else {
			i += 1
			j = i + 1
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
