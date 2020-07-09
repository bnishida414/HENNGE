package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
)

func main() {
	// Initialize the scanner
	scanner := bufio.NewScanner(os.Stdin)
	// Take in the first line of input
	scanner.Scan()

	// N is the number of test cases to follow
	N := scanner.Text()
	fmt.Println("Test cases:", N)

	// We don't have access to for statements,
	// so we'll have to use recursion.

}

func testCaseLoop() {

}

func integerLoop() {

}
