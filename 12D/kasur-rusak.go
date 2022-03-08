package main

import (
	"fmt"
	"strings"
)

var s string

func palindrom(s string) string {
	if len(s) <= 1 {
		return "YA"
	}
	if s[:1] == s[len(s)-1:] {
		s = strings.TrimPrefix(s, s[:1])
		s = strings.TrimSuffix(s, s[len(s)-1:])
		return palindrom(s)
	}
	return "BUKAN"
}

func main() {
	fmt.Scan(&s)
	fmt.Print(palindrom(s))
}
