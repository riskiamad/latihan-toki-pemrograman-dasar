package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s, res string
	var split []string
	fmt.Scan(&s)
	if strings.Contains(s, "_") {
		split = strings.Split(s, "_")
		for id, word := range split {
			if id != 0 {
				word = strings.Title(word)
			}
			res += word
		}
	} else {
		for _, letter := range s {
			i := strconv.QuoteRune(letter)
			if i == strings.ToUpper(i) {
				res += "_"
				i = strings.ToLower(i)
			}
			res += i
		}
	}

	fmt.Print(res)
}
