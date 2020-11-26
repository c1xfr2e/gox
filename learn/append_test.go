package learn

import (
	"fmt"
	"testing"
)

func Append(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) {
		newSlice := make([]byte, l+len(data), 2*(l+len(data)))
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	return slice
}

func TestAppend(t *testing.T) {
	a := []byte("123")
	b := Append(a, []byte("456"))
	fmt.Printf("len: %d  cap:%d\n", len(b), cap(b))
	fmt.Println(string(b))
}

func TestReSlice(t *testing.T) {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	c := b[1:4]
	fmt.Println(string(c))
	c[0] = 'x'
	fmt.Println(string(b))

	s := []int{0, 1, 2, 3, 4, 5}
	s0 := s[2:5]
	fmt.Println(len(s0), cap(s0))
	s1 := s0[:cap(s0)]
	fmt.Println((s1))
	fmt.Println(len(s1), cap(s1))
}
