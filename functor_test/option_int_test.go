package functor_test

import (
	"fmt"

	. "github.com/sh4nnongoh/goFunctional/functor"
)

func ExampleOptionalIntFunctor_Map() {
	const val = 5
	e := SomeInt(val).Map(plusOne).Map(minusOne)
	if e.Int() != val {
		fmt.Println("Map Error")
	}
	// Output:
}

func ExampleEmptyInt() {
	e := EmptyInt()
	if !e.Empty() {
		fmt.Println("Not Empty")
	}
	e.Map(plusOne).Map(minusOne)
	if !e.Empty() {
		fmt.Println("Not Empty")
	}
	// Output:
}

func ExampleSomeInt() {
	e := SomeInt(5)
	if e.Empty() {
		fmt.Println("Empty")
	}
	e.Map(plusOne).Map(minusOne)
	if e.Empty() {
		fmt.Println("Empty")
	}
	// Output:
}
