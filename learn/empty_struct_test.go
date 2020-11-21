package learn

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestSizeof(t *testing.T) {
	var i int
	var s string
	var c complex128
	var arr [2]int
	var sli []int
	var st struct{}
	fmt.Println(unsafe.Sizeof(i))   //
	fmt.Println(unsafe.Sizeof(s))   //
	fmt.Println(unsafe.Sizeof(c))   //
	fmt.Println(unsafe.Sizeof(arr)) //
	fmt.Println(unsafe.Sizeof(sli)) //
	fmt.Println(unsafe.Sizeof(st))  //
}

// go:linkname zerobase runtime.zerobase
var zerobase struct{}

func TestEmptyStruct(t *testing.T) {
	var a, b struct{}
	fmt.Println(&a == &b) // false
	fmt.Println(&a == &zerobase)

	var c, d struct{}
	fmt.Printf("%p %p\n", &c, &d)
	fmt.Println(&c == &d) // true
	fmt.Println(&c == &zerobase)
}

func TestEmptyStruct2(t *testing.T) {

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
