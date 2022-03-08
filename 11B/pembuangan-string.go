package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	for i := 0; i < 100; i++ {
		s = strings.ReplaceAll(s, t, "")
		if !strings.Contains(s, t) {
			break
		}
	}

	fmt.Print(s)
}
