package main

import "fmt"

func main() {
	var m, n, p, x, totalMulti int
	fmt.Scan(&m, &n, &p)
	var matriksA = make([][]int, m)
	var matriksB = make([][]int, n)
	var matriksC = make([][]int, m)
	for i := 0; i < m; i++ {
		matriksA[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&x)
			matriksA[i][j] = x
		}
	}
	for i := 0; i < n; i++ {
		matriksB[i] = make([]int, p)
		for j := 0; j < p; j++ {
			fmt.Scan(&x)
			matriksB[i][j] = x
		}
	}
	for i := 0; i < m; i++ {
		matriksC[i] = make([]int, p)
		for j := 0; j < p; j++ {
			for k := 0; k < n; k++ {
				totalMulti = totalMulti + (matriksA[i][k] * matriksB[k][j])
			}
			matriksC[i][j] = totalMulti
			totalMulti = 0
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < p; j++ {
			fmt.Printf("%d ", matriksC[i][j])
		}
		fmt.Print("\n")
	}
}
