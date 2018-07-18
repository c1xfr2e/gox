package algorithm

import (
	"fmt"
)

func lower(A []int, x int) int {
	a, b := 0, len(A)
	for a < b {
		m := (a + b) / 2
		if A[m] >= x {
			b = m
		} else {
			a = m + 1
		}
	}
	return b
}

func upper(A []int, x int) int {
	a, b := 0, len(A)
	for a < b {
		m := (a + b) / 2
		if A[m] > x {
			b = m
		} else {
			a = m + 1
		}
	}
	return b
}

func TestBounds() {
	A := []int{1, 2, 3, 3, 3, 4, 5, 6, 7}
	fmt.Println(lower(A, 3))
	fmt.Println(upper(A, 3))
}
