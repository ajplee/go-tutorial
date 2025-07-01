package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func inOrder(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	inOrder(t.Left, ch)
	ch <- t.Value
	inOrder(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	inOrder(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if v1 != v2 {
			return false
		}

		if !ok1 || !ok2 {
			break
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(300), tree.New(300)))
	fmt.Println(Same(tree.New(1), tree.New(200)))
}
