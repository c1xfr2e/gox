package algorithm

import (
	"sort"
	"testing"
)

func radixSort(a []int, maxBase int) {
	c := make([]int, 10+1)
	b := make([]int, len(a))
	for base := 1; base < maxBase; base *= 10 {
		for i := range c {
			c[i] = 0
		}
		for _, elem := range a {
			d := (elem / base) % 10
			c[d+1]++
		}
		for i := 1; i < len(c); i++ {
			c[i] += c[i-1]
		}
		for _, elem := range a {
			d := (elem / base) % 10
			b[c[d]] = elem
			c[d]++
		}
		for i, elem := range b {
			a[i] = elem
		}
	}
}

func Test_radixSort(t *testing.T) {

	cases := []struct {
		arr  []int
		base int
	}{
		{[]int{5, 3, 6, 1, 2, 7, 4, 8}, 10},
		{[]int{19, 18, 17, 16, 15, 14, 13, 12, 11, 10}, 100},
		{[]int{123, 156, 231, 333, 422, 123, 831}, 1000},
	}
	for _, c := range cases {
		radixSort(c.arr, c.base)
		if !sort.IntsAreSorted(c.arr) {
			t.Errorf("failed on case: %v", c)
		}
	}
}
