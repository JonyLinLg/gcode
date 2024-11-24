package gpool

import "sync"

type iTask interface {
	Run()
	Free()
	Next() iTask
	SetNext(n iTask)
}

type task struct {
	f    func()
	next iTask
}

var taskPool sync.Pool

func init() {
	taskPool = sync.Pool{New: func() any { return &task{} }}
}

func getTask(f func()) iTask {
	t := taskPool.Get().(*task)
	t.f = f
	return t
}

func (i *task) Run() {
	i.f()
}

func (i *task) Free() {
	i.f = nil
	i.next = nil
	taskPool.Put(i)
}

func (i *task) Next() iTask {
	return i.next
}

func (i *task) SetNext(n iTask) {
	i.next = n
}

// iTaskList is a linked list of tasks
type iTaskList interface {
	Push(t iTask)
	Pop() iTask
	IsEmpty() bool
	Len() int
}

// taskList is a linked list of tasks
type taskList struct {
	head iTask
	tail iTask
	l    int
}

// newTaskList returns a new taskList
func newTaskList() iTaskList {
	return &taskList{}
}

// Len returns the length of the taskList
func (i *taskList) Len() int {
	return i.l
}

func (i *taskList) Push(t iTask) {
	if i.head == nil {
		i.head = t
		i.tail = t
	} else {
		i.tail.SetNext(t)
		i.tail = t
	}
	i.l++
}

func (i *taskList) Pop() iTask {
	if i.head == nil {
		return nil
	}
	t := i.head
	i.head = t.Next()
	i.l--
	t.SetNext(nil)
	if i.head == nil {
		i.tail = nil
	}
	return t
}

func (i *taskList) IsEmpty() bool {
	return i.head == nil
}
