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
	var result int
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		numTrim := strings.TrimSpace(input)
		num, _ := strconv.Atoi(numTrim)
		result += num
	}
	fmt.Print(result)
}
