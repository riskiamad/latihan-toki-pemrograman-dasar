package main

import "fmt"

func main() {
	var alas, tinggi float64

	fmt.Scan(&alas, &tinggi)
	luas := (alas * tinggi) / 2
	fmt.Printf("%.2f", luas)
}
