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

func fac(n int) int {
	chs := make([]chan int, n+1)
	for i := 0; i < n+1; i++ {
		chs[i] = make(chan int)
	}
	for i := 0; i < n; i++ {
		go func(i int) {
			for n := range chs[i] {
				fmt.Printf("channel[%d] read %d\n", i, n)
				if n == 1 {
					chs[i] <- 1
					break
				}
				if n > 1 {
					chs[i+1] <- n - 1
					r := <-chs[i+1]
					chs[i] <- n * r
				}
			}
		}(i)
	}
	chs[0] <- n
	return <-chs[0]
}

func Test_fac(t *testing.T) {
	fmt.Println(fac(5))
}
