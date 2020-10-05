package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	isFirst := true
	prevNum := 0
	currNum := 1
	return func() int {
		if isFirst {
			// first Fib number is 0
			isFirst = false
			return 0
		}
		result := currNum + prevNum
		prevNum = currNum
		currNum = result
		return result
	}
}

func runEx4() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
