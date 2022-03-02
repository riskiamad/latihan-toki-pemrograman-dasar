package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)
	floor := math.Floor(n)
	ceiling := math.Ceil(n)
	fmt.Print(floor, ceiling)
}
