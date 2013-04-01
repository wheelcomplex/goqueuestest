// implements locking queues, using list and mutex
package goqueuestest

type pcLifoElement struct {
	value interface{}
	prev *pcLifoElement
}

type PcLifo struct {
	head *pcLifoElement
	lock chan int
}

func NewPcLifo() *PcLifo {
	q := &PcLifo{}
	q.head = nil
	q.lock = make(chan int, 1)
	q.lock <- 1
	return q
}

func (q *PcLifo) Enqueue(value interface{}) {
	<-q.lock
	q.head = &pcLifoElement{value, q.head}
	q.lock <- 1
}

func (q *PcLifo) Dequeue() (interface{}, bool) {
	<-q.lock
	if q.head == nil {
		q.lock <- 1
		return nil, false
	}
	value := q.head.value
	q.head = q.head.prev
	q.lock <- 1
	return value, true
}

type pcFifoElement struct {
	value interface{}
	next *pcFifoElement
}

type PcFifo struct {
	front *pcFifoElement
	back *pcFifoElement
	lock   chan int
}

func NewPcFifo() *PcFifo {
	q := &PcFifo{}
	q.lock = make(chan int, 1)
	q.lock <- 1
	q.front = nil
	q.back = nil
	return q
}

func (q *PcFifo) Enqueue(value interface{}) {
	<-q.lock
	tmp := &pcFifoElement{value, nil}
	if q.front == nil {
		q.front = tmp
		q.back = tmp
	} else {
		q.back.next = tmp
		q.back = tmp
	}
	q.lock <- 1
}

func (q *PcFifo) Dequeue() (value interface{}, bool) {
	<-q.lock
	tmp := q.front
	if q.front != nil {
		q.front = q.front.next
	} else {
		q.back = nil
	}
	q.lock <- 1
	if tmp != nil {
		return tmp.value, true
	}
	return nil, false
}
