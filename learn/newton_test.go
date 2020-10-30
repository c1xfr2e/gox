package learn

import (
	"fmt"
	"testing"
)

func TestNewtonSqrt(t *testing.T) {
	x := 256.0
	p := 1.0
	q := 0.0
	PREC := 0.000000000000001
	for !(p-q > -PREC && p-q < PREC) {
		q, p = p, p-(p*p-x)/(p*2)
	}
	fmt.Println(p)
}
