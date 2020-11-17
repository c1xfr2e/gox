package csp

import (
	"fmt"
	"math"
	"testing"
)

// EratosthenesPlain returns prime numbers <= n in a slice.
func EratosthenesPlain(n int) []int {
	// phi is the count of prime numbers <=n approximated by x/ln(x).
	phi := int(math.Ceil((float64(n) / math.Log(float64(n)) * 1.2)))
	ps := make([]int, phi)
	ps[0], ps[1] = 2, 3
	sqrt := int(math.Sqrt(float64(n)))
	k := 2
	for m := 3; m <= n; m += 2 {
		yes := true
		for i := 0; i < k; i++ {
			if ps[i] > sqrt {
				break
			}
			if m%ps[i] == 0 {
				yes = false
				break
			}
		}
		if yes {
			ps[k] = m
			k++
		}
	}
	return ps[:k]
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
