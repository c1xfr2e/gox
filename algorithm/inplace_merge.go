package algorithm

import "fmt"

func lower_bound(A []int, left int, right int, x int) int {
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

func upper_bound(A []int, left int, right int, x int) int {
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

	first_cut, second_cut := 0, 0
	if true || len1 > len2 {
		first_cut = first + (mid-first)/2
		second_cut = lower_bound(A, mid, last, A[first_cut])
	} else {
		second_cut = mid + (last-mid)/2
		first_cut = upper_bound(A, first, mid, A[second_cut])
	}

	Rotate(A, first_cut, mid, second_cut)

	new_mid := mid
	if second_cut-mid > 0 {
		new_mid = first_cut + (second_cut - mid)
	}

	InplaceMerge(A, first, first_cut, new_mid)
	InplaceMerge(A, new_mid, second_cut, last)
}

func TestInplaceMerge() {
	A := []int{1, 2, 2, 2, 2}
	InplaceMerge(A, 0, 1, len(A))
	fmt.Println(A)

	A = []int{1, 2, 3, 3, 3, 4, 5, 6, 7}
	fmt.Println(lower_bound(A, 3, len(A), 3))
	fmt.Println(upper_bound(A, 3, len(A), 3))
}
