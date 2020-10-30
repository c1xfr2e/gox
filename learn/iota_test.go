package learn

import (
	"fmt"
	"testing"
)

const (
	A = iota + 1
	B
	C = 5
	D = 99
)

const (
	E float64 = iota + 1
	F
	G int32 = 10
	H
	I
)

func TestIota(t *testing.T) {
	fmt.Println(A, B, C, D)
	fmt.Println(E, F, G, H, I)
	fmt.Printf("%T, %T, %T, %T, %T \n", E, F, G, H, I)
}
