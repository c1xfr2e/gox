package learn

import (
	"fmt"
	"testing"
	"time"
)

const IntervalPeriod time.Duration = 3 * time.Second

type myTimer struct {
	timer *time.Timer
}

func (t *myTimer) update() {
	nextTick := time.Now().Add(IntervalPeriod)
	fmt.Println(nextTick, "- next tick")
	diff := nextTick.Sub(time.Now())
	if t.timer == nil {
		t.timer = time.NewTimer(diff)
	} else {
		t.timer.Reset(diff)
	}
}

func TestTimer(t *testing.T) {
	tm := &myTimer{}
	tm.update()
	for {
		<-tm.timer.C
		fmt.Println(time.Now(), "- just ticked")
		tm.update()
	}
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1000 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(2600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
