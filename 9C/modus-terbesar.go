package main

import "fmt"

func main() {
	var numbers = []int{}
	var n, x, count, savedCount, modus int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		numbers = append(numbers, x)
	}
	len := len(numbers)
	for i := 0; i < len; i++ {
		if numbers[i] > 0 {
			for j := i + 1; j < len; j++ {
				if numbers[i] == numbers[j] {
					numbers[j] = 0
					count += 1
				}
			}

			if count > savedCount {
				savedCount = count
				modus = numbers[i]
			} else if count == savedCount {
				if numbers[i] > modus {
					savedCount = count
					modus = numbers[i]
				}
			}
		}
		count = 0
	}
	fmt.Print(modus)
}
