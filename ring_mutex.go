package goqueuestest

import (
	"container/ring"
	"sync"
)

type RLifo struct {
	head     *ring.Ring // last enqueued value
	length   int
	capacity int
	m        sync.Mutex
}

func NewRLifo() *RLifo {
	q := &RLifo{}
	q.length = 0
	q.capacity = growBy
	q.head = ring.New(growBy)
	q.m = sync.Mutex{}
	return q
}

func (q *RLifo) Enqueue(value interface{}) {
	q.m.Lock()
	if q.length >= q.capacity {
		q.capacity = q.capacity + growBy
		q.head.Link(ring.New(growBy))
	}
	q.head = q.head.Next()
	q.head.Value = value
	q.length += 1
	q.m.Unlock()
}

func (q *RLifo) Dequeue() (value interface{}, ok bool) {
	q.m.Lock()
	if q.length == 0 {
		q.m.Unlock()
		return nil, false
	}
	value = q.head.Value
	q.head = q.head.Prev()
	q.length -= 1
	q.m.Unlock()
	return value, true
}

type RFifo struct {
	head     *ring.Ring // last enqueued value
	tail     *ring.Ring // last dequeued value
	length   int
	capacity int
	m        sync.Mutex
}

func NewRFifo() *RFifo {
	q := &RFifo{}
	q.length = 0
	q.capacity = growBy
	q.head = ring.New(growBy)
	q.tail = q.head
	q.m = sync.Mutex{}
	return q
}

func (q *RFifo) Enqueue(value interface{}) {
	q.m.Lock()
	if q.length >= q.capacity {
		q.capacity = q.capacity + growBy
		q.head.Link(ring.New(growBy))
	}
	q.head = q.head.Next()
	q.head.Value = value
	q.length += 1
	q.m.Unlock()
}

func (q *RFifo) Dequeue() (value interface{}, ok bool) {
	q.m.Lock()
	if q.length == 0 {
		q.m.Unlock()
		return nil, false
	}
	q.tail = q.tail.Next()
	value = q.tail.Value
	q.length -= 1
	q.m.Unlock()
	return value, true
}
