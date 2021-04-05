package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	)


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Type Smthing")

	scanner.Scan()

	input, err := scanner.Text()

	strconv.ParseInt(input, 10, 64)
	fmt.Printf("You typed %q", input)

}