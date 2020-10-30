package learn

import (
	"fmt"
	"log"
	"net"
	"testing"
)

const listenAddr = "localhost:4000"

func TestListenAccept(t *testing.T) {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(c, "Hello")
		c.Close()
	}
}
