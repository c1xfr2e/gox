package learn

import (
	"fmt"
	"testing"
)

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type S struct {
	A string
	B int
}

func TestInterface(t *testing.T) {
	describe(S{A: "hello", B: 123})
	describe(&S{A: "hello", B: 123})
}

// Staff to test partial implementation of an interface.
type Intr interface {
	MethodA()
	MethodB()
}

type ImplA struct {
	Foo string
}

func (m *ImplA) MethodA() {
	fmt.Printf("ImplA.MethodA: %s\n", m.Foo)
}

type ImplB struct {
	Bar string
}

func (m *ImplB) MethodB() {
	fmt.Printf("ImplB.MethodB: %s\n", m.Bar)
}

type ImplAB struct {
	ImplA
	ImplB
}

func TestPartialImplementation(t *testing.T) {
	m := &ImplAB{
		ImplA: ImplA{Foo: "FOO"},
		ImplB: ImplB{Bar: "BAR"},
	}
	// Assign m to variable of Intr.
	var i Intr
	i = m
	i.MethodA()
	i.MethodB()
}

func convertIntr(tp IntrType, val interface{}) string {
	switch tp {
	case typeA:
		a := val.(*ImplA)
		return a.Foo
	case typeB:
		b := val.(*ImplB)
		return b.Bar
	default:
		return "unknown"
	}
}

type IntrType int

const (
	typeA IntrType = iota
	typeB
)

func TestConvertIntr(t *testing.T) {
	fmt.Println(convertIntr(typeA, &ImplA{"HI A"}))
	fmt.Println(convertIntr(typeB, &ImplB{"HI B"}))
	// will panic
	fmt.Println(convertIntr(typeA, &ImplB{"panic"}))
}
