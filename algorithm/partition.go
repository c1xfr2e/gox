package algorithm

import "fmt"

// PartitionITA implements the partition algorithm in "Introduction to Algoritm 3E"
func PartitionITA(A []int) int {
	r := len(A) - 1
	x := A[r]
	p := -1
	for j := p + 1; j < r; j++ {
		if A[j] <= x {
			p++
			A[p], A[j] = A[j], A[p]
		}
	}
	A[p+1], A[r] = A[r], A[p+1]
	return p + 1
}

// PartitionLeftRight implements partitioning by scaning elements from left and right.
func PartitionLeftRight(A []int) int {
	r := len(A) - 1
	x := A[r]
	i, j := 0, r-1
	for {
		if i >= j {
			fmt.Printf("i >= j (%d >= %d) at beginning of loop\n", i, j)
			fmt.Println(A)
			break
		}
		for A[i] < x {
			i++
		}
		for j > 0 && A[j] > x {
			j--
		}
		if i >= j {
			break
		}
		A[i], A[j] = A[j], A[i]
		i++
		j--
	}
	A[i], A[r] = A[r], A[i]
	return i
}

// Partition3Way implements 3-way partitioning.
// Return when elements in A[lo, hi] equal the pivot.
func Partition3Way(A []int) (lo, hi int) {
	a, c := -1, len(A)-1
	x := A[len(A)-1]
	for i := 0; i < c; {
		switch {
		case A[i] == x:
			{
				i++
			}
		case A[i] < x:
			{
				a++
				A[a], A[i] = A[i], A[a]
				i++
			}
		case A[i] > x:
			{
				c--
				A[i], A[c] = A[c], A[i]
			}
		}
	}
	A[c], A[len(A)-1] = A[len(A)-1], A[c]
	return a + 1, c
}
