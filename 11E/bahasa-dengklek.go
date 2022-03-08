package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s, res string
	fmt.Scan(&s)

	for _, word := range s {
		i := strconv.QuoteRune(word)
		if i == strings.ToUpper(i) {
			res += strings.ToLower(i)
		} else {
			res += strings.ToUpper(i)
		}
	}
	fmt.Print(string(res))

}
