package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Important notes:
// 1. No output until all input has been received
// 2. No for statements
// 3. Negative integers are ignored

// To handle note 1: here is our slice to store
// our values before sending to output
var sumSquares []int

func main() {
	// Initialize the scanner
	scanner := bufio.NewScanner(os.Stdin)
	// Take in the first line of input
	scanner.Scan()

	// N is the number of test cases to follow
	N, _ := strconv.Atoi(scanner.Text())
	fmt.Println("Test cases:", N)

	// We don't have access to for statements,
	// so we'll have to use recursion.
	testCaseLoop(N)
}

// Loop through each test case
func testCaseLoop(N int) {
	// Check base case
	if N <= 0 {
		return
	}
	// Save the sum of squares for this test case

	// Recursion!
	testCaseLoop(N - 1)
}

// Loop through each integer and return its square
func integerLoop(X int) {
	// Check base case
	if X == 0 {
		return
	}
	if 
	// Recursion!
	integerLoop(X - 1)
}