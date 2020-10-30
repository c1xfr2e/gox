package learn

import (
	"golang.org/x/tour/tree"
)

func walkImpl(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkImpl(t.Left, ch)
	ch <- t.Value
	walkImpl(t.Right, ch)
}

// Walk walks the tree t sending all values from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkImpl(t, ch)
	close(ch)
}

// Same determines whether the trees t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	go Walk(t1, c1)
	c2 := make(chan int)
	go Walk(t2, c2)
	for {
		n1, ok1 := <-c1
		n2, ok2 := <-c2
		if !ok1 && !ok2 {
			return true
		}
		if ok1 != ok2 || n1 != n2 {
			return false
		}
	}
}
