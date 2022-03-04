package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var result string
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		result += line
	}
	fmt.Print(result)
}
