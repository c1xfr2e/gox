package algorithm

import "fmt"

func sink(i, n int, A []int) {
	li := 2*i + 1
	for li < n {
		ri := li + 1
		next := li
		if ri < n && A[li] < A[ri] {
			next = ri
		}
		if A[i] < A[next] {
			A[i], A[next] = A[next], A[i]
			i = next
			li = 2*i + 1
		} else {
			break
		}
	}
}

func heap_sort(A []int) {
	n := len(A)
	k := n / 2
	for i := k - 1; i >= 0; i-- {
		sink(i, n, A)
	}

	for i := n; i > 1; i-- {
		A[0], A[i-1] = A[i-1], A[0]
		sink(0, i-1, A)
	}
}

func isHeap(A []int, top int) bool {
	left := 2*top + 1
	leftisheap := left >= len(A) || isHeap(A, left)
	right := 2*top + 2
	rightisheap := right >= len(A) || isHeap(A, right)
	return leftisheap && rightisheap

	li := 2*top + 1
	ri := 2*top + 2
	if li < len(A) {
		if !isHeap(A, li) {
			return false
		}
		if A[top] < A[li] {
			return false
		}
	}
	if ri < len(A) {
		if !isHeap(A, ri) {
			return false
		}
		if A[top] < A[ri] {
			return false
		}
	}

	return true
}

func partition(A []int, left int, right int) int {
	x := A[right]
	p := left - 1
	for j := p + 1; j < right; j++ {
		if A[j] <= x {
			p++
			A[p], A[j] = A[j], A[p]
		}
	}
	A[p+1], A[right] = A[right], A[p+1]
	return p + 1
}

func quicksort(A []int, left int, right int) {
	if left >= right {
		return
	}
	mid := partition(A, left, right)
	quicksort(A, left, mid-1)
	quicksort(A, mid+1, right)
}

func TestHeap() {
	b := []int{5, 3, 1, 6, 9, 0, 7, 8, 4, 2}
	quicksort(b, 0, len(b)-1)
	fmt.Println(b)

	fmt.Println(isHeap([]int{1, 2, 3}, 0))
	fmt.Println(isHeap([]int{3, 2, 1}, 0))

	a := []int{5, 3, 1, 6, 9, 0, 7, 8, 4, 2}
	fmt.Println(a)
	fmt.Println(isHeap(a, 0))
	heap_sort(a)
	fmt.Println(a)
	//fmt.Println(is_heap(a, 0))
}
