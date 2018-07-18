package algorithm

import "fmt"

func get_next(s string) []int {
	N := make([]int, len(s))
	N[0] = -1
	for i := 1; i < len(N); i++ {
		p := i - 1
		for ; N[p] != -1; p = N[p] {
			if s[i-1] == s[N[p]] {
				break
			}
		}
		N[i] = N[p] + 1
	}
	return N
}

func kmp_searh(s, p string) int {
	next := get_next(p)
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
				i += 1
				j = 0
			}
		}
	}

	if j == M {
		return i - M
	} else {
		return -1
	}
}

func TestKMP() {
	T := "ababababaabaaabaab"
	P := "ababaab"
	i := kmp_searh(T, P)
	if i >= 0 {
		fmt.Println(i, T[i:i+len(P)])
	} else {
		fmt.Println(P, "not found")
	}
}
