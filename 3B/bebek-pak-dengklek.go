package main

import "fmt"

func main() {

	var bebekJantan, bebekBetina int

	fmt.Scan(&bebekJantan)
	fmt.Scan(&bebekBetina)
	totalBebek := bebekJantan + bebekBetina
	fmt.Println(totalBebek)
}
