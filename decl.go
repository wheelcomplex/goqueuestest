package goqueuestest

import "unsafe"

type Queue interface {
	Enqueue(value interface{})
	Dequeue() (value interface{}, ok bool)
}

// common for lockfree things

type Element struct {
	Value interface{}
	Next  unsafe.Pointer
}

const growBy = 1000