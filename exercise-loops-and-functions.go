package main

import (
	"fmt"
	"math"
)

// improveGuess create new guess according to Newton's formula
func improveGuess(currGuess, x float64) float64 {
	return currGuess - (currGuess*currGuess-x)/(2*currGuess)
}

// SqrtUsing10Iterations Calculating the sqrt of x using Newton's method
// f(z) = z^2 - x
// f'(z) = 2z
// X_n+1 = X_n - f(X_n)/f'(X_n)
func SqrtUsing10Iterations(x float64) float64 {
	const numIterations int = 10
	var result float64 = 1.0 // initial guess
	for i := 0; i < numIterations; i++ {
		result = improveGuess(result, x)
	}
	return result
}

// SqrtUsingEpsilon Calculates sqrt of x using Newton's method
// Stops when the change in calculated values is smaller than epsilon
func SqrtUsingEpsilon(x float64) (int, float64) {
	const epsilon float64 = 0.00000000001
	var result float64 = 1.0
	var newResult float64 = improveGuess(result, x)
	var iteration int = 1 // as we already did 1 calculation
	for math.Abs(newResult-result) > epsilon {
		result = newResult
		newResult = improveGuess(result, x)
		iteration++
	}
	result = newResult
	return iteration, result
}

func m1ain() {
	var x float64 = math.Pow(16719, 2)
	fmt.Printf("Ater 10 iterations: %v\n", SqrtUsing10Iterations(x))
	numIterations, calculatedValue := SqrtUsingEpsilon(x)
	fmt.Printf("After %v iterations: %v", numIterations, calculatedValue)
}
