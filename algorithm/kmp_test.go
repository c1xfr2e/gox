package algorithm

import (
	"fmt"
	"testing"
)

func genNext(s string) []int {
	next := make([]int, len(s))
	next[0] = -1
	for i := 1; i < len(next); i++ {
		p := i - 1
		// search util a match if found or p is 0 (next[p] == -1)
		for next[p] != -1 {
			if s[i-1] == s[next[p]] {
				break
			}
			p = next[p]
		}
		next[i] = next[p] + 1
	}
	return next
}

func KmpSearch(s, p string) int {
	next := genNext(p)
	N, M := len(s), len(p)
	i, j := 0, 0
	for i < N && j < M {
		if s[i] == p[j] {
			i, j = i+1, j+1
			continue
		} else {
			if next[j] > 0 {
				j = next[j]
			} else {
				i++
				j = 0
			}
		}
	}
	if j == M {
		return i - M
	}
	return -1
}

func TestKMP(t *testing.T) {
	T := "ababababaabaaabaab"
	P := "ababaab"
	i := KmpSearch(T, P)
	if i >= 0 {
		fmt.Println(i, T[i:i+len(P)])
	} else {
		fmt.Println(P, "not found")
	}
}
