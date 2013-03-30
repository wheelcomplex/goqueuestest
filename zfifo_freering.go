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
	// Creating an initial node
	node := unsafe.Pointer(&node_t{nil, unsafe.Pointer(q)})

	// Both head and tail point to the initial node
	q.head = node
	q.tail = node

	q.length = 0
	q.capacity = 1000
	q.freelist = ring.New(q.capacity)
	q.m = sync.Mutex{}
	return q
}

func (q *ZFifoFreering) newNode() *node_t {
	q.m.Lock()
	if q.length == 0 {
		q.m.Unlock()
		return &node_t{}
	}
	value, _ := (q.freelist.Value).(*node_t)
	q.freelist = q.freelist.Prev()
	q.m.Unlock()
	return value
}

func (q *ZFifoFreering) freeNode(elem *node_t) {
	q.m.Lock()
	if q.length >= q.capacity {
		q.m.Unlock()
		return
	}
	q.freelist = q.freelist.Next()
	q.freelist.Value = elem
	q.m.Unlock()
}

// Enqueue inserts the value at the tail of the queue
func (q *ZFifoFreering) Enqueue(value interface{}) {
	node := q.newNode()
	//new(node_t) // Allocate a new node from the free list
	node.value = value // Copy enqueued value into node
	node.next = unsafe.Pointer(q)
	for { // Keep trying until Enqueue is done
		tail := atomic.LoadPointer(&q.tail)

		// Try to link in new node
		if atomic.CompareAndSwapPointer(&(*node_t)(tail).next, unsafe.Pointer(q), unsafe.Pointer(node)) {
			// Enqueue is done.  Try to swing tail to the inserted node.
			atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(node))
			return
		}

		// Try to swing tail to the next node as the tail was not pointing to the last node
		atomic.CompareAndSwapPointer(&q.tail, tail, (*node_t)(tail).next)
	}
}

// Dequeue returns the value at the head of the queue and true, or if the queue is empty, it returns a nil value and false
func (q *ZFifoFreering) Dequeue() (value interface{}, ok bool) {
	for {
		head := atomic.LoadPointer(&q.head)               // Read head pointer
		tail := atomic.LoadPointer(&q.tail)               // Read tail pointer
		next := atomic.LoadPointer(&(*node_t)(head).next) // Read head.next
		if head != q.head {                               // Check head, tail, and next consistency
			continue // Not consistent. Try again
		}

		if head == tail { // Is queue empty or tail failing behind
			if next == unsafe.Pointer(q) { // Is queue empty?
				return
			}
			// Try to swing tail to the next node as the tail was not pointing to the last node
			atomic.CompareAndSwapPointer(&q.tail, tail, next)
		} else {
			// Read value before CAS
			// Otherwise, another dequeue might free the next node
			value = (*node_t)(next).value
			// Try to swing Head to the next node
			if atomic.CompareAndSwapPointer(&q.head, head, next) {
				ok = true
				q.freeNode((*node_t)(head))
				return
			}
			value = nil
		}
	}
	return // Dummy return
}
