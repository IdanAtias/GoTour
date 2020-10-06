package main

import (
	"fmt"
)

// ErrNegativeSqrt type definition
type ErrNegativeSqrt float64

// ErrNegativeSqrt implements error interface
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt function
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	_, sqrtResult := SqrtUsingEpsilon(x)
	return sqrtResult, nil
}

func runEx6() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
