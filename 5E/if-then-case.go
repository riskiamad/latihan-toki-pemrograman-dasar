package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	strN := strconv.Itoa(n)
	if len(strN) == 1 {
		fmt.Print("satuan")
	} else if len(strN) == 2 {
		fmt.Print("puluhan")
	} else if len(strN) == 3 {
		fmt.Print("ratusan")
	} else if len(strN) == 4 {
		fmt.Print("ribuan")
	} else if len(strN) == 5 {
		fmt.Print("puluhribuan")
	} else if len(strN) == 6 {
		fmt.Print("ratusribuan")
	}

}
