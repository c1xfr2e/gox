package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func workFn(n int64) func() int {
	return func() int {
		fmt.Println(n)
		return int(n)
	}
}

// Request is.
type Request struct {
	fn func() int // The operation to perform.
	c  chan int   // The channel to return the result.
}

func requester(work chan<- Request) {
	c := make(chan int)
	for {
		// Kill some time (fake load).
		n := rand.Int63n(5000)
		time.Sleep(time.Millisecond * time.Duration(n))
		work <- Request{workFn(n), c} // send request
		result := <-c                 // wait for answer
		fmt.Println(result)
	}
}

// Worker is.
type Worker struct {
	requests chan Request // work to do (buffered channel)
	pending  int          // count of pending tasks
	index    int          // index in the heap
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests // get Request from balancer
		req.c <- req.fn()   // call fn and send result
		done <- w           // we've finished this request
	}
}

// Pool is the heap array to keep all workers.
type Pool []*Worker

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p Pool) Len() int {
	return len(p)
}

// Push into heap.
func (p Pool) Push(i interface{}) {
	p = append(p, i.(*Worker))
}

// Pop from heap
func (p Pool) Pop() interface{} {
	w := p[len(p)-1]
	p = p[0 : len(p)-1]
	return w
}

func (p Pool) Swap(i, j int) {
	w := p[i]
	p[i] = p[j]
	p[j] = w
}

// Balancer is.
type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work: // received a Request...
			b.dispatch(req) // ...so send it to a Worker
		case w := <-b.done: // a worker has finished ...
			b.completed(w) // ...so update its info
		}
	}
}

// Send Request to worker
func (b *Balancer) dispatch(req Request) {
	// Grab the least loaded worker...
	w := heap.Pop(&b.pool).(*Worker)
	// ...send it the task.
	w.requests <- req
	// One more in its work queue.
	w.pending++
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

// Job is complete; update heap
func (b *Balancer) completed(w *Worker) {
	// One fewer in the queue.
	w.pending--
	// Remove it from heap.
	heap.Remove(&b.pool, w.index)
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

func main() {
	// we'll have 15 workers
	numWorkers := 15
	doneChan := make(chan *Worker, numWorkers)

	// create a balancer
	b := Balancer{
		pool: make([]*Worker, 0, numWorkers),
		done: doneChan,
	}

	// generate the workers
	for i := 0; i < numWorkers; i++ {
		w := &Worker{
			requests: make(chan Request, 50),
			pending:  0,
			index:    i,
		}
		b.pool = append(b.pool, w)
		// go to work
		go w.work(doneChan)
	}

	workChan := make(chan Request)
	go requester(workChan)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		b.balance(workChan)
		wg.Done()
	}()
	wg.Wait()
}
