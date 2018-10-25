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

func ExampleIntSliceFunctor_Ints() {
	ints := []int{0, 1, 2, 3, 4}
	functor := LiftIntSlice(ints)
	functorInts := functor.Ints()
	for i := range ints {
		if ints[i] != functorInts[i] {
			fmt.Println("Not Equal")
			return
		}
	}
	fmt.Println("Equal")
	// Output: Equal
}

func ExampleIntSliceFunctor_String() {
	ints := []int{0, 1, 2, 3, 4}
	functor := LiftIntSlice(ints)
	fmt.Println(functor.String())
	// Output: []int{0, 1, 2, 3, 4}
}
