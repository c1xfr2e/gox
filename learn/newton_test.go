package learn

import (
	"fmt"
	"testing"
)

func TestNewtonSqrt(t *testing.T) {
	c := 256.0
	x := 1.0
	p := 0.0
	PREC := 0.000000000000001
	for !(x-p > -PREC && x-p < PREC) {
		p, x = x, x-(x*x-c)/(x*2)
	}
	fmt.Println(x)
}
