package algorithm

import (
	"testing"
)

func maxUniqueSubstrBrute(str string) string {
	var first, size int
	for i := 0; i < len(str); i++ {
		m := make(map[byte]bool)
		j := i
		for ; j < len(str); j++ {
			if m[str[j]] {
				break
			}
			m[str[j]] = true
		}
		if j-i > size {
			size = j - i
			first = i
		}
	}
	return str[first : first+size]
}

func maxUniqueSubstr(str string) string {
	var start, size int
	for i := 0; i < len(str); {
		m := make(map[byte]int)
		next := -1
		j := i
		for j < len(str) {
			b := str[j]
			if p, has := m[b]; has {
				next = p + 1
				break
			}
			m[b] = j
			j++
			next = j
		}
		if j-i > size {
			size = j - i
			start = i
		}
		i = j
		if next != -1 {
			i = next
		}
	}
	return str[start : start+size]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxUniqueSubstrLinear(str string) string {
	lastIndice := make(map[byte]int)
	first := 0
	length := 0
	i := 0
	for j := 0; j < len(str); j++ {
		li, ok := lastIndice[str[j]]
		lastIndice[str[j]] = j
		if ok && li >= i {
			i = li + 1
			continue
		}
		if j-i+1 > length {
			length = j - i + 1
			first = i
		}
	}
	return str[first : first+length]
}

func Test_maxUniqueSubstr(t *testing.T) {
	cases := []struct{ str, sub string }{
		{"aaa", "a"},
		{"abab", "ab"},
		{"aaaxa", "ax"},
		{"abca123dd", "bca123d"},
		{"ekurqq98273sd", "q98273sd"},
	}
	for _, c := range cases {
		r := maxUniqueSubstr(c.str)
		if r != c.sub {
			t.Errorf("failed on case %q: want %q, got %q", c.str, c.sub, r)
		}
	}
}

func Test_maxUniqueSubstrBrute(t *testing.T) {
	cases := []string{
		"3284578569902317",
		"jekljgaks384aksdjf",
		"helworoadknni",
		"你暗红色2我就开打时",
		"ddfkdskfje9349wsklajf",
	}
	for _, c := range cases {
		r1 := maxUniqueSubstrBrute(c)
		r2 := maxUniqueSubstr(c)
		if r1 != r2 {
			t.Errorf("failed on case %q: want %q, got %q", c, r1, r2)
		}
	}
}

func Test_maxUniqueSubstrLinear(t *testing.T) {
	cases := []string{
		"xxx",
		"xyxy",
		"3284578569902317",
		"jekljgaks384aksdjf",
		"helworoadknni",
		"你暗红色2我就开打时",
		"ddfkdskfje9349wsklajf",
	}
	for _, c := range cases {
		r1 := maxUniqueSubstrLinear(c)
		r2 := maxUniqueSubstr(c)
		if r1 != r2 {
			t.Errorf("failed on case %q: want %q, got %q", c, r1, r2)
		}
	}

}
