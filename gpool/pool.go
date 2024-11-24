package gpool

import (
	"sync"
	"sync/atomic"
)

type iPool interface {
	Go(f func())
}

var defaultPool = newPool()

func Go(f func()) {
	defaultPool.Go(f)
}

type pool struct {
	lock  sync.Mutex
	tasks iTaskList

	workers atomic.Int32
}

func newPool() iPool {
	return &pool{
		tasks: newTaskList(),
	}
}

func (p *pool) Go(f func()) {
	p.lock.Lock()
	t := getTask(f)
	p.tasks.Push(t)
	p.lock.Unlock()
	if p.workers.Load() > 3 {
		return
	}
	w := newWorker(p)
	w.Run()
}

func (p *pool) GetTask() iTask {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.tasks.Pop()
}

func (p *pool) AddWorker() {
	p.workers.Add(1)
}

func (p *pool) RemoveWorker() {
	p.workers.Add(-1)
}
