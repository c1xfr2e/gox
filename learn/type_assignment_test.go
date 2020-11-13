package learn

import (
	"fmt"
	"testing"
)

type MyString = string

func TestTypeAssignment(t *testing.T) {
	var mt MyString
	mt = "hello"
	fmt.Printf("%T %v\n", mt, mt)
}
