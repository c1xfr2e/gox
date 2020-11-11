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
}

// Heapify makes A a heap.
func Heapify(A []int) {
	if len(A) < 2 {
		return
	}
	for i := len(A)/2 - 1; i >= 0; i-- {
		shiftdown(A, i)
	}
}

func shiftdown(A []int, top int) {
	pos := top
	temp := A[pos]
	child := pos*2 + 1
	for child < len(A) {
		rightchild := child + 1
		if rightchild < len(A) && A[rightchild] < A[child] {
			child = rightchild
		}
		A[pos] = A[child]
		pos = child
		child = 2*pos + 1
	}
	A[pos] = temp
	shiftup(A, pos, top)
}

func shiftup(A []int, i int, top int) {
	temp := A[i]
	for i > top {
		up := (i - 1) >> 1
		if temp < A[up] {
			A[i] = A[up]
			i = up
		} else {
			break
		}
	}
	A[i] = temp
}

func checkHeap(A []int, top int) bool {
	left := 2*top + 1
	checkLeft := left >= len(A) || (A[top] <= A[left] && isHeap(A, left))
	right := 2*top + 2
	checkRight := right >= len(A) || (A[top] <= A[right] && isHeap(A, right))
	return checkLeft && checkRight
}

func TestHeap() {
	fmt.Println(isHeap([]int{1, 2, 3}, 0))
	fmt.Println(isHeap([]int{3, 2, 1}, 0))

	a := []int{5, 3, 1, 6, 9, 0, 7, 8, 4, 2}
	fmt.Println(a)
	fmt.Println(isHeap(a, 0))
	heap_sort(a)
	fmt.Println(a)
	//fmt.Println(is_heap(a, 0))
}
