package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walk func(t *tree.Tree, ch chan int)
	// walk := func...だと関数内関数でバインドされない。
	walk = func(t *tree.Tree, ch chan int) {
		if t == nil {
			return
		}
		walk(t.Left, ch)
		ch <- t.Value
		walk(t.Right, ch)
	}

	walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for val := range ch1 {
		if val != <-ch2 {
			return false
		}
	}

	return true
}

func testWalk() {
	fmt.Println("start Walk Test")
	n := 1
	ch := make(chan int, 10)
	go Walk(tree.New(n), ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("end Walk Test")
}

func testSame() {
	fmt.Println("start Walk Test")
	fmt.Println(Same(tree.New(1), tree.New(1)) == true)
	fmt.Println(Same(tree.New(1), tree.New(2)) == false)
	fmt.Println("end Walk Test")
}

func main() {
	testWalk()
	testSame()
}
