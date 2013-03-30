package zq

import (
	"sync/atomic"
	"unsafe"
)

// based on http://www.cs.cmu.edu/~410-s05/lectures/L31_LockFree.pdf
// Warning! Suffers from ABA problem

type ZLifo struct {
	head unsafe.Pointer
}

func NewZLifo() *ZLifo {
	return &ZLifo{nil}
}

func (q *ZLifo) Enqueue(value interface{}) {
	node := unsafe.Pointer(&Element{value, nil})
	for {
		(*Element)(node).Next = q.head
		if atomic.CompareAndSwapPointer(&q.head, (*Element)(node).Next, node) {
			break
		}
	}
}

func (q *ZLifo) Dequeue() (value interface{}, ok bool) {
	current := q.head
	for current != nil {
		if atomic.CompareAndSwapPointer(&q.head, current, (*Element)(current).Next) {
			value = (*Element)(current).Value
			return value, true
		}
		current = q.head
	}
	return nil, false
}
