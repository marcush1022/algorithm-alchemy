// package golang_demo
// name tracker

package golang_demo

import (
	"context"
	"fmt"
	"time"
)

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{ch: make(chan string, 10)}
}

func (t *Tracker) Run() {
	fmt.Println(">>>>>>>> Run")
	// deal with one data per second
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(">>>>>>>>> data", data)
	}
	// stop tracker
	t.stop <- struct{}{}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	fmt.Println(">>>>>>>> Event", data)
	select {
	case t.ch <- data: // save 10 data
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Shutdown(ctx context.Context) error {
	fmt.Println(">>>>>>>> Shutdown")
	close(t.ch)
	select {
	case <-t.stop: // if tracker stop
		break
	case <-ctx.Done(): // if timeout
		return ctx.Err()
	}
	return nil
}

func TrackerMain() {
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(), "aaa")
	_ = tr.Event(context.Background(), "bbb")
	_ = tr.Event(context.Background(), "ccc")

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	if err := tr.Shutdown(ctx); err != nil {
		fmt.Println(">>>>>>>>>>>>>>> Shutdown Err", err)
	}
}

func TrackerMain2() {
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(), "aaa")
	_ = tr.Event(context.Background(), "bbb")
	_ = tr.Event(context.Background(), "ccc")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := tr.Shutdown(ctx); err != nil {
		fmt.Println(">>>>>>>>>>>>>>> Shutdown Err", err)
	}
}
