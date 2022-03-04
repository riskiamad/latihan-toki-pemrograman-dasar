package main

import "fmt"

func main() {
	var m, n, x int
	fmt.Scan(&m, &n)
	var numbers = make([][]int, m)
	var rotate = make([][]int, m)

	for i := 0; i < m; i++ {
		numbers[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&x)
			numbers[i][j] = x
		}
	}

	for i := 0; i < m; i++ {
		rotate[i] = make([]int, n)
		for j := 0; j < n; j++ {
			rotate[i][j] = numbers[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := m - 1; j >= 0; j-- {
			fmt.Printf("%d ", rotate[j][i])
		}
		fmt.Print("\n")
	}

}
