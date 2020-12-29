package algorithm

import (
	"math/rand"
	"sort"
	"testing"
)

func qsort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	p := a[n-1]
	i := -1
	for j := 0; j < n-1; j++ {
		if a[j] < p {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[n-1] = a[n-1], a[i+1]
	qsort(a[:i+1])
	qsort(a[i+2:])
}

func Test_qsort(t *testing.T) {
	cases := [][]int{
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		[]int{5, 3, 6, 1, 2, 7, 4, 8},
		[]int{1, 1, 1, 1},
		[]int{2, 2, 1, 1},
		[]int{1, 1, 2, 2},
		func() []int {
			ra := [10]int{}
			for i := range ra {
				ra[i] = rand.Intn(20)
			}
			return ra[:]
		}(),
	}
	for _, c := range cases {
		qsort(c)
		if !sort.IntsAreSorted(c) {
			t.Errorf("failed on case: %v", c)
		}
	}
}
