package algorithm

import (
	"fmt"
	"testing"
)

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func knapsack01(W []int, V []int, cap int) int {
	n := len(W) - 1
	m := make([][]int, n+1)
	for i := range m {
		m[i] = make([]int, cap+1)
	}
	for i := range m[0] {
		m[0][i] = 0
	}
	for i := 1; i <= n; i++ {
		m[i][0] = 0
		for j := 1; j <= cap; j++ {
			if j >= W[i] {
				m[i][j] = maxInt(m[i-1][j], m[i-1][j-W[i]]+V[i])
			} else {
				m[i][j] = m[i-1][j]
			}
		}
	}
	return m[n][cap]
}

func TestKnapsack01(t *testing.T) {
	W := []int{-1, 23, 26, 20, 18, 32, 27, 29, 26, 30, 27}
	V := []int{-1, 505, 353, 458, 220, 354, 414, 498, 545, 473, 543}
	fmt.Println(knapsack01(W, V, 67))
}
