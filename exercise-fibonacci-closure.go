package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// 0 1 1 2 3 5 8 13 21 34 ...
func fibonacci() func() int {
	currNum := 0
	nextNum := 1
	return func() int {
		result := currNum
		currNum, nextNum = nextNum, result+nextNum
		return result
	}
}

func runEx4() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
