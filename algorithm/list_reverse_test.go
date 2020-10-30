package algorithm

import (
	"fmt"
	"testing"
)

type list struct {
	value int
	next  *list
}

// reverse the order of nodes in l and return the new list.
// eg. from `A->B->C` to `C->B->A`
func (l *list) reverse() *list {
	var p *list // p tracks current head node
	for l != nil {
		q := l.next
		l.next = p
		p = l
		l = q
	}
	return p
}

func TestListReverse(t *testing.T) {
	l := &list{
		1, &list{2, &list{3, &list{4, nil}}},
	}
	h := l.reverse()
	for i := h; i != nil; i = i.next {
		fmt.Println(i.value)
	}
}
