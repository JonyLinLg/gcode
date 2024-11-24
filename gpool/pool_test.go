package gpool

import (
	"testing"
	"time"
)

func fool(i int) {
	time.Sleep(time.Millisecond * 100)
	println("hello", i)
}

func TestPool(t *testing.T) {
	p := newPool()
	for i := 0; i < 100; i++ {
		p.Go(func() {
			fool(i)
		})
	}

	time.Sleep(time.Second * 5)
}
