package csp

import (
	"fmt"
	"testing"
)

// 4.2 Recursion: Factorial
//
// > "Problem: Compute a factorial by the recursive method, to a given
// limit."

func fac(p chan int, q chan int, r chan int, o chan int) {
	for n := range p {
		if n == 1 {
			o <- 1
			break
		}
		if n > 1 {
			q <- n - 1
			o <- n * <-r
		}
	}
}

func fac0(n int) int {
	p1 := make(chan int)
	q1 := make(chan int)
	r1 := make(chan int)
	o1 := make(chan int)
	go fac(p1, q1, r1, o1)

	p2 := q1
	q2 := make(chan int)
	r2 := make(chan int)
	o2 := r1
	go fac(p2, q2, r2, o2)

	p3 := q2
	q3 := make(chan int)
	r3 := make(chan int)
	o3 := r2
	go fac(p3, q3, r3, o3)

	p4 := q3
	o4 := r3
	go fac(p4, nil, nil, o4)

	p1 <- n
	return <-o1
}

func Test_fac0(t *testing.T) {
	fmt.Println(fac0(4))
}
