package algorithm

import (
	"testing"
)

// match represents a common substring match
type match struct {
	i1, i2 int // index
	len    int // length
	s      string
}

func doLCS(s1, s2 string) match {
	l1, l2 := len(s1), len(s2)
	m := make([][]int, l1)
	for i := range m {
		m[i] = make([]int, l2)
	}
	var r match
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			var p int
			if i == 0 || j == 0 {
				p = -1
			} else {
				p = m[i-1][j-1]
			}
			if s1[i] == s2[j] {
				l := p + 1
				m[i][j] = l
				if m[i][j] > r.len {
					r = match{
						i1:  i - l + 1,
						i2:  j - l + 1,
						len: l,
						s:   s1[i-l+1 : i+1],
					}
				}
			} else {
				m[i][j] = 0
			}
		}
	}
	return r
}

func Test_doLCS(t *testing.T) {
	cases := []struct {
		s1, s2 string
		m      match
	}{
		{
			"abc",
			"123",
			match{},
		},
		{
			"helloxworld",
			"xllworlx",
			match{6, 3, 4, "worl"},
		},
		{
			"xabc1234ijkxyz",
			"234abc12ijkx",
			match{1, 3, 5, "abc12"},
		},
	}
	for _, c := range cases {
		m := doLCS(c.s1, c.s2)
		if m != c.m {
			t.Errorf("failed on case %v", c)
		}
	}
}
