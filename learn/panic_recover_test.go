package learn

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestPanicRecover(t *testing.T) {
	{
		s, err := first()
		fmt.Printf("play: Returned from first: err=%v\n", err)
		fmt.Printf("s=%v\n", s)
	}
}

var errPanic = errors.New("Cumstom Panic Error")

func first() (s string, err error) {
	// defer 1
	defer fmt.Println("first: defer 1")

	// defer 2 with recover
	defer func() {
		e := recover()
		if panicErr, ok := e.(error); ok && panicErr == errPanic {
			fmt.Println("first: defer 2: set err after recover")
			err = panicErr
			s = "recovered"
		}
		fmt.Printf("first: defer 2: recovered in first: %v\n", e)
	}()

	// defer 3
	defer fmt.Println("first: defer 3")

	fmt.Println("first: before calling second.")
	r := second()
	fmt.Println("first: after calling second.", r)

	return "final", nil
}

func second() string {
	fmt.Println("second: before panic")
	panic(errPanic)
	fmt.Println("second: after panic")
	return "second"
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
