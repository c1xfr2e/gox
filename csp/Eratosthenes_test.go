package csp

import (
	"fmt"
	"math"
	"testing"
)

// Eratosthenes returns prime numbers <= n in a slice.
// phi is the count of prime numbers approximate to x/ln(x).
func Eratosthenes(n int) []int {
	phi := int(math.Ceil((float64(n) / math.Log(float64(n)) * 1.2)))
	p := make([]int, phi)
	p[0], p[1] = 2, 3
	k := 2
	sq := int(math.Sqrt(float64(n)))
	for m := 3; m <= n; m += 2 {
		yes := true
		for i := 0; i < k; i++ {
			if p[i] > sq {
				break
			}
			if m%p[i] == 0 {
				yes = false
				break
			}
		}
		if yes {
			p[k] = m
			k++
		}
	}
	return p[:k]
}

func TestEratosthenes(t *testing.T) {
	e := Eratosthenes(100000)
	fmt.Println(len(e))
}
