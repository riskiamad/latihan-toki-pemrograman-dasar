package main

import "fmt"

func main() {
	var x, y, reversedDigitX, reversedDigitY, reversedDigitZ int
	fmt.Scan(&x, &y)

	for x > 0 {
		lastDigit := x % 10
		reversedDigitX = reversedDigitX*10 + lastDigit

		x = x / 10
	}

	for y > 0 {
		lastDigit := y % 10
		reversedDigitY = reversedDigitY*10 + lastDigit

		y = y / 10
	}
	z := reversedDigitY + reversedDigitX

	for z > 0 {
		lastDigit := z % 10
		reversedDigitZ = reversedDigitZ*10 + lastDigit

		z = z / 10
	}
	fmt.Print(reversedDigitZ)
}
