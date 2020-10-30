package learn

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	var m1 map[string]string
	if m1 == nil {
		fmt.Println("m1 is nil")
	}
	if val, ok := m1["key"]; ok {
		fmt.Printf("m1[\"key\"]=%q\n", val)
	} else {
		fmt.Println("key not in m1")
	}

	m2 := map[string]bool{}
	m2["a"] = true
	m2["b"] = true
	fmt.Println(m2["c"])
	if m2["a"] {
		fmt.Println("a in m2")
	}
	if m2["d"] {
		fmt.Println("d in m2")
	}
}
