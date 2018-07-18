package algorithm

import "fmt"

func merge_sort_r(A []int, T []int, left int, right int) {
	if left >= right {
		return
	}

	for i := left; i <= right; i++ {
		T[i] = A[i]
	}

	mid := (left + right) / 2
	merge_sort_r(T, A, left, mid)
	merge_sort_r(T, A, mid+1, right)

	i, j, k := left, mid+1, left
	for ; i <= mid && j <= right; k++ {
		if T[i] < T[j] {
			A[k] = T[i]
			i++
		} else {
			A[k] = T[j]
			j++
		}
	}
	for ; i <= mid; i, k = i+1, k+1 {
		A[k] = T[i]
	}
	for ; j <= right; j, k = j+1, k+1 {
		A[k] = T[j]
	}
}

func merge_sort(A []int) {
	N := len(A)
	tmp := make([]int, N)
	merge_sort_r(A, tmp, 0, N-1)
}

func TestMergeSort() {
	A := []int{5, 3, 1, 6, 9, 0, 7, 8, 4, 2}
	merge_sort(A)
	fmt.Println(A)
}
