package timeoutctx

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func printSlice(ctx context.Context, sl []int) {
	for _, x := range sl {
		select {
		case <-ctx.Done():
			fmt.Println("ctx Done")
			return
		default:
			fmt.Println(x)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func TestTimeoutContext(t *testing.T) {
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		printSlice(ctx, sl[0:len(sl)/2])
	}()
	go func() {
		defer wg.Done()
		printSlice(ctx, sl[len(sl)/2:])
	}()

	c := make(chan error)
	go func() {
		defer close(c)
		wg.Wait()
	}()

	select {
	case <-c:
		fmt.Println("finished")
	case <-ctx.Done():
		fmt.Println("timeout")
	}

	time.Sleep(time.Second)
}
