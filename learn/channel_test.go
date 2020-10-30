package learn

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	c := make(chan int)
	c <- 12
	fmt.Println(<-c)
}
