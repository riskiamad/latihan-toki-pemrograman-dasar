package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1, s2, s3, s4, res string
	fmt.Scan(&s1, &s2, &s3, &s4)
	res1 := strings.Split(s1, s2)[0]
	res2 := strings.Split(s1, s2)[1]
	res = res1 + res2
	res1 = strings.Split(res, s3)[0]
	res2 = strings.Split(res, s3)[1]
	res = res1 + s3 + s4 + res2
	fmt.Println(res)
}
