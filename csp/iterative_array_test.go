package csp

import (
	"fmt"
	"testing"
)

func addBy1(n int) int {
	chs := make([]chan int, n)
	chs[0] = make(chan int)
	for i := 1; i < n; i++ {
		chs[i] = make(chan int)
	}
	for i := 1; i < n; i++ {
		go func(i int) {
			chs[i] <- <-chs[i-1] + 1
		}(i)
	}
	chs[0] <- 1
	r := <-chs[n-1]
	for _, c := range chs {
		close(c)
	}
	return r
}

func Test_addBy1(t *testing.T) {
	fmt.Println(addBy1(100))
}
