package learn

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// Play plays.
func Play() {
	{
		s, i, a, err := f(2)
		fmt.Printf("play: Returned from f: err=%v\n", err)
		fmt.Printf("s=%v\n", s)
		fmt.Printf("i=%v\n", i)
		fmt.Printf("a=%v\n", a)
	}
	{
		s, i, a, err := f(3)
		fmt.Printf("play: Returned from f: err=%v\n", err)
		fmt.Printf("s=%v\n", s)
		fmt.Printf("i=%v\n", i)
		fmt.Printf("a=%v\n", a)
	}
}

func f(n int) (s string, i int, a [3]int, err error) {
	// defer 1
	defer fmt.Println("f: defer 1")

	// defer 2 with recover
	defer func() {
		e := recover()
		if panicErr, ok := e.(error); ok && panicErr == errPanic {
			err = panicErr
			s = "recovered"
		}
		fmt.Printf("f: defer 2: Recovered in f: %v\n", e)
	}()

	// defer 3
	defer fmt.Println("f: defer 3")

	fmt.Println("f: before calling g.")
	r := g(n)
	fmt.Println("f: returned normally from g.", r)

	return "fifteen", 15, a, nil
}

var errPanic = errors.New("Cumstom Panic Error")

func g(i int) [3]int {
	if i == 3 {
		fmt.Println("g: before panic 3")
		panic(errPanic)
	}
	if i == 2 {
		fmt.Println("g: before panic 2")
		panic(fmt.Sprintf("panic parameter: i=%v", i))
	}
	defer fmt.Println("g: Defer ", i)
	fmt.Println("g: Printing ", i)
	fmt.Println(g(i + 1))

	return [3]int{1, 2, 3}
}

func TestPlay(t *testing.T) {
	Play()
}

func TestDeferInfiniteLoop(t *testing.T) {
	tm := time.NewTimer(5 * time.Second)
L:
	for {
		defer func() {
			fmt.Println("defer in TestDeferInfiniteLoop")
		}()
		select {
		case <-tm.C:
			break L
		default:
			fmt.Println("default")
			time.Sleep(time.Second)
		}
	}
	fmt.Println("exit TestDeferInfiniteLoop")
}
