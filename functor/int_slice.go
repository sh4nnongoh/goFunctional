package functor

import (
	"fmt"
	"sync"
)

// IntSliceFunctor returns a container of []int, and acts an interface for eaasily iterating over a give slice of ints,
// applying a given function on each element, and returning the new IntSliceFunctor with the new results.
//
// All implementations of this interrface must adherre to the follwowing rules:
//
// 1. f.Map(func(i int){return i}) == f
//		-> mapping with an identify function does nothing
// 2. f.Map(funcA(funcB(param))) == f.Map(funcA).Map(funcB)
//		-> functions can be composed and executed in serial
type IntSliceFunctor interface {
	fmt.Stringer
	// Map is the Functor's function. It applies the fn to every element in the contained slice.
	Map(fn func(int) int) IntSliceFunctor
	// Ints is just a convenience function to get the int slice that the functor holds.
	Ints() []int
}

type intSliceFunctorImpl struct {
	ints []int
}

// LiftIntSlice convets an int slice into an IntSLiceFunctor.
// In FP terms, this operation is called "lifting".
func LiftIntSlice(slice []int) IntSliceFunctor {
	return intSliceFunctorImpl{ints: slice}
}

// Map executes fn on every int in isf's internal slice and returns the resultant ints
func (isf intSliceFunctorImpl) Map(fn func(int) int) IntSliceFunctor {
	if len(isf.ints) < 100 {
		isf.ints = serialIntMapper(isf.ints, fn)
		return isf
	}
	isf.ints = parallelIntMapper(isf.ints, fn)
	return isf
}

// Ints just returns a copy of the ints in isf
func (isf intSliceFunctorImpl) Ints() []int {
	return isf.ints
}

// String just returns the sting representation of the functor.
func (isf intSliceFunctorImpl) String() string {
	return fmt.Sprintf("%#v", isf.ints)
}

// Serial Constructs
func serialIntMapper(ints []int, fn func(int) int) []int {
	for i, elt := range ints {
		retInt := fn(elt)
		ints[i] = retInt
	}
	return ints
}

// Parallel Constructs
type parallelIntMapperResult struct {
	idx int
	val int
}

func parallelIntMapper(ints []int, fn func(int) int) []int {
	ch := make(chan parallelIntMapperResult)
	var wg sync.WaitGroup
	for i, elt := range ints {
		wg.Add(1)
		go func(idx int, elt int) {
			defer wg.Done()
			ch <- parallelIntMapperResult{idx: idx, val: fn(elt)}
		}(i, elt)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for elt := range ch {
		ints[elt.idx] = elt.val
	}
	return ints
}
