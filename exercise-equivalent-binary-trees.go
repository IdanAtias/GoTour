package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// we'll traverse the tree in-order. that is- left son, father, right son.
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func runEx10() {
	t1 := tree.New(1)
	fmt.Println(t1)
	ch := make(chan int, 10)
	go Walk(t1, ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	t2 := tree.New(1)
	t3 := tree.New(2)
	if !Same(t1, t2) || Same(t1, t3) || Same(t2, t3) {
		panic("FAIL!")
	}
	fmt.Println("PASS!")
}
