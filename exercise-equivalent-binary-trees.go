package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func () {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func () {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
        n1, ok1 := <-ch1
        n2, ok2 := <-ch2

        if n1 != n2 || ok1 != ok2 {
            return false
        }

        if !ok1 {
            break
        }
    }

    return true

}

func main() {
	fmt.Println(Same(tree.New(1),tree.New(1)))
}

