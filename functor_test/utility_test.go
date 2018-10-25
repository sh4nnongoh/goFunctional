package functor_test

import (
	"fmt"
)

// Test Helper Functions
func intSlice(numInts int) []int {
	slice := make([]int, numInts)
	for i := 0; i < numInts; i++ {
		slice[i] = i + 5
	}
	return slice
}

func plusOne(i int) int {
	return i + 1
}

func minusOne(i int) int {
	return i - 1
}

func testEq(a, b []int) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Example Tests Functions
func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}
