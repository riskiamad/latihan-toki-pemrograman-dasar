package main

import (
	"fmt"
	"strings"
)

var value string
var n int

func caesar(r rune, shift int) rune {
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}
func main() {
	fmt.Scan(&value, &n)
	value2 := strings.Map(func(r rune) rune {
		return caesar(r, n)
	}, value)
	fmt.Print(value2)
}
