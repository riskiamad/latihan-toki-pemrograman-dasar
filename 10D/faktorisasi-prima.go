package main

import "fmt"

func main() {
	var n, a, divide int
	var fakPrim = []int{}
	var prime bool
	fmt.Scan(&n)
	for j := 1; j <= n; j++ {
		if divide > 2 {
			prime = false
			divide = 0
			break
		} else {
			if n%j == 0 {
				divide += 1
			}
			prime = true
		}
	}

	if !prime {
		for i := 2; i <= n; i++ {
			if n%i == 0 {
				fakPrim = append(fakPrim, i)
				n = n / i
				i = 1
			}
		}

		for id, fak := range fakPrim {
			if id != len(fakPrim)-1 {
				if fak == fakPrim[id+1] {
					a += 1
				} else {
					if a != 0 {
						fmt.Printf("%d^%d x ", fak, a+1)
					} else {
						fmt.Printf("%d x ", fak)
					}
					a = 0
				}
			} else {
				if fak == fakPrim[id-1] {
					fmt.Printf("%d^%d", fak, a+1)
				} else {
					fmt.Print(fak)
				}
			}

		}
	} else {
		fmt.Print(n)
	}
}
