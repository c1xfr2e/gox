package csp

import (
	"fmt"
	"math"
	"testing"
)

// EratosthenesPlain returns prime numbers <= n in a slice.
// phi is the count of prime numbers approximate to x/ln(x).
func EratosthenesPlain(n int) []int {
	phi := int(math.Ceil((float64(n) / math.Log(float64(n)) * 1.2)))
	p := make([]int, phi)
	p[0], p[1] = 2, 3
	k := 2
	sq := int(math.Sqrt(float64(n)))
	for m := 3; m <= n; m += 2 {
		yes := true
		for i := 0; i < k; i++ {
			if p[i] > sq {
				break
			}
			if m%p[i] == 0 {
				yes = false
				break
			}
		}
		if yes {
			p[k] = m
			k++
		}
	}
	return p[:k]
}

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- int) {
	ch <- 2
	for i := 3; ; i += 2 {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes.
func EratosthenesConcurrent() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	print(<-ch, "\n")
	print(<-ch, "\n")
}

func TestEratosthenesPlain(t *testing.T) {
	e := EratosthenesPlain(100000)
	fmt.Println(len(e))
}

func TestEratosthenesConcurrent(t *testing.T) {
	EratosthenesConcurrent()
}
