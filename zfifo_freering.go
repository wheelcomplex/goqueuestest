package goqueuestest

import (
	"container/ring"
	"sync"
	"sync/atomic"
	"unsafe"
)

type ZFifoFreering struct {
	head     unsafe.Pointer
	tail     unsafe.Pointer
	freelist *ring.Ring
	length   int
	capacity int
	m        sync.Mutex
}

func NewZFifoFreering() *ZFifoFreering {
	q := &ZFifoFreering{}
	empty := unsafe.Pointer(&Element{})
	q.head = empty
	q.tail = empty
	q.length = 0
	q.capacity = 1000
	q.freelist = ring.New(q.capacity)
	q.m = sync.Mutex{}
	return q
}

func (q *ZFifoFreering) newElem() *Element {
	q.m.Lock()
	if q.length == 0 {
		q.m.Unlock()
		return &Element{}
	}
	value := (q.freelist.Value).(*Element)
	q.freelist = q.freelist.Prev()
	q.m.Unlock()
	return value
}

func (q *ZFifoFreering) freeElem(elem *Element) {
	q.m.Lock()
	if q.length >= q.capacity {
		q.m.Unlock()
		return
	}
	q.freelist = q.freelist.Next()
	q.freelist.Value = elem
	q.m.Unlock()
}

func (q *ZFifoFreering) Enqueue(value interface{}) {
	// Allocate a new node
	var tail, next, node unsafe.Pointer
	element := q.newElem()
	element.Value = value
	element.Next = nil
	node = unsafe.Pointer(element)

	for { // Keep trying until Enqueue is done
		tail = q.tail                  // Read Tail.ptr and Tail.count together
		next = ((*Element)(tail)).Next // Read next ptr and count fields together
		if tail == q.tail {            // Are tail and next consistent?

			if next == nil { // Was Tail pointing to the last node?
				// Try to link node at the end of the linked list
				if atomic.CompareAndSwapPointer(&((*Element)(tail)).Next, next, node) {
					break // Enqueue is done.  Exit loop
				}
			} else { // Tail was not pointing to the last node
				atomic.CompareAndSwapPointer(&q.tail, tail, next) // Try to swing Tail to the next node
			}
		}
	}
	// Enqueue is done.  Try to swing Tail to the inserted node
	atomic.CompareAndSwapPointer(&q.tail, tail, node)
}

func (q *ZFifoFreering) Dequeue() (value interface{}, ok bool) {
	var head, tail, next unsafe.Pointer
	for {
		head = q.head                  // Read Head
		tail = q.tail                  // Read Tail
		next = ((*Element)(head)).Next // Read Head.ptr->next
		if head == q.head {            // Are head, tail, and next consistent?
			if head == tail {
				if next == nil {
					return nil, false // Queue is empty, couldn't dequeue
				}
				// Tail is falling behind.  Try to advance it
				atomic.CompareAndSwapPointer(&q.tail, tail, next)
			} else {
				// Read value before CAS
				// Otherwise, another dequeue might free the next node
				value = ((*Element)(next)).Value
				// Try to swing Head to the next node
				if atomic.CompareAndSwapPointer(&q.head, head, next) {
					break // Dequeue is done.  Exit loop
				}
			}
		}
	}
	// It is safe now to free the old node
	q.freeElem((*Element)(next))
	return value, true
}
