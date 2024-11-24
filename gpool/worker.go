package gpool

import (
	"fmt"
	"sync"
)

type iWorker interface {
	Run()
}

var workerPool sync.Pool

func init() {
	workerPool.New = func() any {
		return &worker{}
	}
}

type worker struct {
	p *pool
}

func newWorker(p *pool) iWorker {
	w := workerPool.Get().(*worker)
	w.p = p
	w.p.AddWorker()
	return w
}

func (w *worker) Run() {
	go func() {
		defer func() {
			workerPool.Put(w)
		}()
		for {
			t := w.p.GetTask()
			if t == nil {
				w.p.RemoveWorker()
				fmt.Println("worker exit")
				break
			}
			t.Run()
			t.Free()
		}
	}()
}
