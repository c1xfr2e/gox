package learn

import (
	"fmt"
	"testing"
)

func TestEmptyStruct(t *testing.T) {
	var a, b struct{}
	fmt.Println(&a == &b) // false
}

func TestEmptyStruct2(t *testing.T) {
	var a, b struct{}
	fmt.Printf("%p %p\n", &a, &b)
	fmt.Println(&a == &b) // true
}

func TestSlice(t *testing.T) {
	// empty struct is GREAT
	a := make([]struct{}, 10)
	b := make([]struct{}, 20)
	fmt.Println(&a == &b)       // false, a and b are different slices
	fmt.Println(&a[0] == &b[0]) // true, their backing arrays are the same

	c := make([]struct{ x int }, 20)
	d := make([]struct{ x int }, 20)
	fmt.Println(&c[0] == &d[0]) // false
}
