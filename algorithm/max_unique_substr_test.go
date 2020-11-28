package algorithm

import (
	"fmt"
	"testing"
)

func maxUniqueSubstr(str string) string {
	var (
		max   int
		index int
		i, j  int
	)
	for i < len(str) {
		m := make(map[byte]int)
		m[str[i]] = 1
		j = i + 1
		for j < len(str) {
			if m[str[j]] == 1 {
				break
			}
			m[str[j]] = 1
			j++
		}
		if j-i > max {
			max = j - i
			index = i
		}
		if j == len(str) {
			break
		}
		for j < len(str) && str[j] == str[j-1] {
			j++
		}
		i = j - 1
	}
	fmt.Println(index, max)
	return str[index : index+max]
}

func Test_maxUniqueSubstr(t *testing.T) {
	fmt.Println(maxUniqueSubstr("abab"))
	fmt.Println(maxUniqueSubstr("aabc1234cc"))
	fmt.Println(maxUniqueSubstr("aaaa"))
}
