package functor_test

import (
	"fmt"

	. "github.com/sh4nnongoh/goFunctional/functor"
)

type intSliceFunctorImpl struct {
	ints []int
}

func ExampleintSliceFunctorImpl() {
	//ints := []int{0, 1, 2, 3, 4}
	//sliceF := intSliceFunctorImpl{ints: slice}
}

func ExampleIntSliceFunctor() {
	ints := []int{0, 1, 2, 3, 4}
	functor := LiftIntSlice(ints)
	fmt.Println(len(functor.Ints()))
	// Output: 5
}

func ExampleIntSliceFunctor_Map() {
	ints := []int{0, 1, 2, 3, 4}
	newInts := []int{0, 1, 2, 3, 4}
	for i := range ints {
		newInts[i] = ints[i]
	}
	functor := LiftIntSlice(newInts)
	mapped := functor.Map(plusOne).Map(minusOne)
	for i := range ints {
		if ints[i] != mapped.Ints()[i] {
			fmt.Println("Not Equal")
			return
		}
	}
	fmt.Println("Equal")
	// Output: Equal
}

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
