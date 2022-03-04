package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numbers = []int{}
	var flippedArray = []int{}
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		numTrim := strings.TrimSpace(input)
		num, _ := strconv.Atoi(numTrim)
		numbers = append(numbers, num)
	}
	len := len(numbers)
	for i := len - 1; i >= 0; i-- {
		flippedArray = append(flippedArray, numbers[i])
	}
	fmt.Println(flippedArray)

}
