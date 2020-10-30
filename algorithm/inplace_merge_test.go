package algorithm

import (
	"fmt"
	"testing"
)

func LowerBound(A []int, left int, right int, x int) int {
	a, b := left, right
	for a < b {
		m := (a + b) / 2
		if A[m] < x {
			a = m + 1
		} else {
			b = m
		}
	}
	return b
}

func UpperBound(A []int, left int, right int, x int) int {
	a, b := left, right
	for a < b {
		m := (a + b) / 2
		if A[m] <= x {
			a = m + 1
		} else {
			b = m
		}
	}
	return b
}

func Reverse(A []int, left int, right int) {
	for i, j := left, right; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}
}

// Rotate A[first,last) toward left, making A[mid] the first element.
func Rotate(A []int, first int, mid int, last int) {
	Reverse(A, first, mid-1)
	Reverse(A, mid, last-1)
	Reverse(A, first, last-1)
}

func InplaceMerge(A []int, first int, mid int, last int) {
	len1, len2 := mid-first, last-mid
	if len1 == 0 || len2 == 0 {
		return
	}
	if len1 == 1 && len2 == 1 {
		if A[mid] < A[first] {
			A[first], A[mid] = A[mid], A[first]
		}
		return
	}

	cut1, cut2 := 0, 0
	// if len1 > len2 {
	cut1 = first + (mid-first)/2
	cut2 = LowerBound(A, mid, last, A[cut1])
	// } else {
	// 	cut2 = mid + (last-mid)/2
	// 	cut1 = UpperBound(A, first, mid, A[cut2])
	// }

	Rotate(A, cut1, mid, cut2)

	newMid := mid
	if cut2-mid > 0 {
		newMid = cut1 + (cut2 - mid)
	}

	InplaceMerge(A, first, cut1, newMid)
	InplaceMerge(A, newMid, cut2, last)
}

func TestInplaceMerge(t *testing.T) {
	A := []int{3, 5, 5, 7, 1, 2, 2, 2, 4, 6}
	InplaceMerge(A, 0, 4, len(A))
	fmt.Println(A)

	A = []int{1, 2, 3, 3, 3, 4, 5, 6, 7}
	fmt.Println(LowerBound(A, 0, len(A), 3), UpperBound(A, 0, len(A), 3))
}
