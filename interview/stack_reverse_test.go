package interview

import (
	"testing"

	"github.com/golang-collections/collections/stack"
)

func TestReverseStack(t *testing.T) {
	wanted := []int{1, 2, 3, 4, 5}
	st := stack.New()
	for _, i := range wanted {
		st.Push(i)
	}
	ReverseStack(st)
	for i := 0; st.Len() > 0; i++ {
		v := st.Pop()
		if v != wanted[i] {
			t.Errorf("Unexpected result: want %d got %d", wanted[i], v)
		}
	}
}

// ReverseStack reverses a stack by Push, Pop and Len methods.
func ReverseStack(st *stack.Stack) {
	if st.Len() == 1 {
		return
	}
	t := st.Pop()
	ReverseStack(st)
	insertAtBottom(st, t)
}

func insertAtBottom(st *stack.Stack, elem interface{}) {
	if st.Len() == 0 {
		st.Push(elem)
		return
	}
	t := st.Pop()
	insertAtBottom(st, elem)
	st.Push(t)
}
