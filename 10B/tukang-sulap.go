package main

import "fmt"

func main() {
	var n, bil, t, x, y int
	var a = []int{}
	var b = []int{}
	var p, q string

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&bil)
		a = append(a, bil)
	}

	for i := 1; i <= n; i++ {
		fmt.Scan(&bil)
		b = append(b, bil)
	}

	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Scan(&p, &x, &q, &y)

		if p == "A" && q == "A" {
			a[x-1], a[y-1] = a[y-1], a[x-1]
		} else if p == "A" && q == "B" {
			a[x-1], b[y-1] = b[y-1], a[x-1]
		} else if p == "B" && q == "B" {
			b[x-1], b[y-1] = b[y-1], b[x-1]
		} else if p == "B" && q == "A" {
			b[x-1], a[y-1] = a[y-1], b[x-1]
		}

	}

	for _, num := range a {
		fmt.Printf("%d ", num)
	}

	fmt.Print("\n")

	for _, num := range b {
		fmt.Printf("%d ", num)
	}
}
