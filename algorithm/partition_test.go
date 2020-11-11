package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPartitionLeftRight(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ar := make([]int, 10)
	for i := range ar {
		ar[i] = rand.Intn(20)
	}
	cases := [][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0},
		[]int{2, 8, 7, 1, 3, 5, 6, 4},
		[]int{7, 6, 5, 4, 3, 2, 1},
		ar,
	}
	for _, ar := range cases {
		m := PartitionLeftRight(ar)
		for i := 0; i < m; i++ {
			if ar[i] > ar[m] {
				t.Fatalf("Failed on case %v", ar)
			}
		}
		for i := m + 1; i < len(ar); i++ {
			if ar[i] < ar[m] {
				t.Fatalf("Failed on case %v", ar)
			}
		}
	}
}

func TestPartition3Way(t *testing.T) {
	rand.Seed(time.Now().Unix())
	ar := make([]int, 30)
	for i := range ar {
		ar[i] = rand.Intn(10)
	}
	cases := [][]int{
		[]int{0, 0, 0, 0},
		[]int{2, 8, 7, 1, 3, 5, 6, 4},
		[]int{5, 3, 1, 3, 2, 8, 6, 3},
		[]int{1, 1, 1, 2, 2, 2},
		[]int{2, 2, 2, 1, 1, 1},
		ar,
	}
	for _, ar := range cases {
		lo, hi := Partition3Way(ar)
		x := ar[lo]
		for i := 0; i < lo; i++ {
			fmt.Printf("%d ", ar[i])
			if !(ar[i] < x) {
				t.Fatalf("Failed on case %v. A[%d] >= %d", ar, i, x)
			}
		}
		for i := lo; i <= hi; i++ {
			fmt.Printf("%d ", ar[i])
			if ar[i] != x {
				t.Fatalf("Failed on case %v. A[%d] != %d", ar, i, x)
			}
		}
		for i := hi + 1; i < len(ar); i++ {
			fmt.Printf("%d ", ar[i])
			if !(ar[i] > x) {
				t.Fatalf("Failed on case %v. A[%d] <= %d", ar, i, x)
			}
		}
		fmt.Printf("[%d,%d]\n", lo, hi)
	}
}
