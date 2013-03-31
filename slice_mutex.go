// implements locking queues, using list and mutex
package goqueuestest

import (
	"sync"
)

type SLifo struct {
	l []interface{}
	last int
	m sync.Mutex
}

func NewSLifo() *SLifo {
	q := &SLifo{}
	q.l = make([]interface{}, 0, growBy)
	q.last = -1
	return q
}

func (q *SLifo) Enqueue(value interface{}) {
	q.m.Lock()
	q.l = append(q.l, value)
	q.last += 1
	q.m.Unlock()
}

func (q *SLifo) Dequeue() (interface{}, bool) {
	q.m.Lock()
	if q.last < 0 {
		q.m.Unlock()
		return nil, false
	}
	value := q.l[q.last]
	q.last -= 1
	q.m.Unlock()
	return value, true
}

type SFifo struct {
	l []interface{}
	tail int
	head int
	length int
	m sync.Mutex
}

func NewSFifo() *SFifo {
	q := &SFifo{}
	q.l = make([]interface{}, growBy)
	q.tail = 0
	q.head = 0
	q.length = 0
	return q
}

func (q *SFifo) Enqueue(value interface{}) {
	q.m.Lock()
	if q.length >= len(q.l) {
		nl := make([]interface{}, len(q.l) + growBy)
		tLen := q.length - q.tail
		copy(nl, q.l[q.tail:])
		copy(nl[tLen:], q.l[:q.head])
		q.l = nl
		q.tail = 0
		q.head = q.length
	}
	q.l[q.head] = value
	q.head = (q.head + 1) % len(q.l)
	q.length += 1
	q.m.Unlock()
}

func (q *SFifo) Dequeue() (interface{}, bool) {
	q.m.Lock()
	if q.length == 0 {
		q.m.Unlock()
		return nil, false
	}
	value := q.l[q.tail]
	q.tail = (q.tail + 1) % len(q.l)
	q.length -= 1
	q.m.Unlock()
	return value, true
}
