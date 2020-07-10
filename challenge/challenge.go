package main

import (
	"fmt"
	"os"
)

// Important notes:
// 1. No output until all input has been received
// 2. No for statements
// 3. Negative integers are ignored

// Handling note 1: use a slice to store our values
// before sending to output
var sumSquares []int

func main() {
	// N is the number of test cases to follow
	var N int
	// Read input N, output error if necessary
	if _, err := fmt.Scanf("%d\n", &N); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// Handling note 2: use recursion instead of
	// for statements
	testCaseLoop(N)

	// Print out the answer
	printArr(N, 0)
}

// Loop through each test case
func testCaseLoop(N int) {
	// Check base case
	if N <= 0 {
		return
	}

	// X is the number of integers to follow
	var X int
	if _, err := fmt.Scanf("%d\n", &X); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// Save sum of squares into slice
	sumSquares = append(sumSquares, integerLoop(X, 0))

	// Next test case!
	testCaseLoop(N - 1)
}

// Add each integer's square to the sum, recursively
func integerLoop(X int, sum int) int {
	// Check base case
	if X <= 0 {
		fmt.Scanln()
		return sum
	}

	// Add square of next integer to sum
	var Yn int
	fmt.Scanf("%d", &Yn)
	// Handling note 3: ignore negatives
	if Yn >= 0 {
		sum += Yn * Yn
	}

	// Next integer!
	return integerLoop(X-1, sum)
}

// Print the sums of squares
func printArr(N int, ct int) {
	if ct >= N {
		return
	}
	fmt.Print(sumSquares[ct])
	// To avoid extra newline at the end
	if ct < N-1 {
		fmt.Println()
	}
	printArr(N, ct+1)
}
