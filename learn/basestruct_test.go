package learn

import (
	"fmt"
	"strings"
	"testing"
)

type BaseSt struct {
	Base1 string
	Base2 int
}

type SubSt struct {
	BaseSt
	Sub1 string
	Sub2 int
	Sub3 float64
}

func convertToBaseSt(i interface{}) error {
	base, ok := i.(BaseSt)
	if !ok {
		return fmt.Errorf("i is not a baseSt")
	}
	fmt.Printf("{%s %d}\n", base.Base1, base.Base2)
	return nil
}

func TestBaseStruct(t *testing.T) {
	ss := []string{"Hello", "world", "!", "12"}
	fmt.Println(strings.Join(ss, ";"))

	sub := SubSt{
		BaseSt: BaseSt{
			Base1: "basestr",
			Base2: 222,
		},
		Sub1: "hello",
		Sub2: 123456,
		Sub3: 3.15,
	}
	err := convertToBaseSt(sub)
	t.Log(err)

	var i interface{}
	i = &sub
	base, ok := i.(*BaseSt)
	if ok {
		fmt.Printf("{%s %d}\n", base.Base1, base.Base2)
	}
}
