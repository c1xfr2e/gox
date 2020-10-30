package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
)

func MergeSortImpl(A []int, B []int, left int, right int) {
	if left >= right {
		return
	}

	// for i := left; i <= right; i++ {
	// 	B[i] = A[i]
	// }

	mid := (left + right) / 2
	MergeSortImpl(B, A, left, mid)
	MergeSortImpl(B, A, mid+1, right)

	i, j, k := left, mid+1, left
	for i <= mid && j <= right {
		if B[i] < B[j] {
			A[k] = B[i]
			k++
			i++
		} else {
			A[k] = B[j]
			k++
			j++
		}
	}
	for ; i <= mid; i, k = i+1, k+1 {
		A[k] = B[i]
	}
	for ; j <= right; j, k = j+1, k+1 {
		A[k] = B[j]
	}
}

func MergeSort(A []int) {
	B := make([]int, len(A))
	copy(B, A)
	MergeSortImpl(A, B, 0, len(A)-1)
}

func TestMergeSort(t *testing.T) {
	a2 := []int{4, 3, 2, 1}
	MergeSort(a2)
	fmt.Println(a2)

	a3 := []int{8, 7, 6, 5, 4, 3, 2, 1}
	MergeSort(a3)
	fmt.Println(a3)

	a1 := []int{2, 5, 3, 4, 1, 6}
	MergeSort(a1)
	fmt.Println(a1)

	// random case
	a4 := make([]int, 20)
	for i := range a4 {
		a4[i] = rand.Intn(30)
	}
	MergeSort(a4)
	fmt.Println(a4)
}
