package learn

import (
	"fmt"
	"testing"
)

type database struct{}

func (db *database) connect() (disconnect func()) {
	fmt.Println("connect")
	return func() {
		fmt.Println("disconnect")
	}
}

func TestDefer(t *testing.T) {
	db := &database{}
	defer db.connect()()

	fmt.Println("query db...")
}
