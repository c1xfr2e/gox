package learn

import (
	"fmt"
	"testing"
)

type Foo struct {
	A int64
	B string
}

type Bar struct {
	*Foo
	B string
	C string
}

func TestEmbeddedStruct(t *testing.T) {
	b := Bar{
		Foo: &Foo{123456, "Foo-B"},
		B:   "Bar-B",
		C:   "Bar-C",
	}
	fmt.Printf("%v\n", b)
	fmt.Println(b.Foo.B, b.B)
}
