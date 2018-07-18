package algorithm

// Heapify is to make A a heap
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
	child := pos * 2 + 1
	for child < len(A) {
		rightchild := child + 1
		if rightchild < len(A) && A[rightchild] < A[child] {
			child = rightchild
		}
		A[pos] = A[child]
		pos = child
		child = 2 * pos + 1
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
	left := 2 * top + 1
	checkLeft := left >= len(A) || (A[top] <= A[left] && isHeap(A, left))
	right := 2 * top + 2
	checkRight := right >= len(A) || (A[top] <= A[right] && isHeap(A, right))
	return checkLeft && checkRight
}
