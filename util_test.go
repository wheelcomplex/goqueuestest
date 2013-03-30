package goqueuestest

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

const (
	testItemCount = 1000
	testTimeout   = 10 * time.Second
)

func timeout(t *testing.T) chan int {
	done := make(chan int)
	go func() {
		select {
		case <-time.After(testTimeout):
			panic("Timeout")
		case <-done:
		}
	}()
	return done
}

func testReturn(t *testing.T, got interface{}, wasOk bool, expected interface{}, expectedOk bool) {
	if wasOk != expectedOk {
		t.Fatalf("Returned ok=%b expected %b", wasOk, expectedOk)
	}

	if got != expected {
		t.Fatalf("Returned %v expected %v", got, expected)
	}
}

func fifoTest(t *testing.T, q Queue) {
	done := timeout(t)

	N := testItemCount
	for i := 0; i < N; i += 1 {
		q.Enqueue(i)
	}

	for i := 0; i < N; i += 1 {
		val, ok := q.Dequeue()
		testReturn(t, val, ok, i, true)
	}

	for i := 0; i < 5; i += 1 {
		val, ok := q.Dequeue()
		testReturn(t, val, ok, nil, false)
	}

	done <- 1
}

func lifoTest(t *testing.T, q Queue) {
	done := timeout(t)

	N := testItemCount
	for i := 0; i < N; i += 1 {
		q.Enqueue(i)
	}

	for i := N - 1; i >= 0; i -= 1 {
		val, ok := q.Dequeue()
		testReturn(t, val, ok, i, true)
	}

	for i := 0; i < 5; i += 1 {
		val, ok := q.Dequeue()
		testReturn(t, val, ok, nil, false)
	}
	done <- 1
}

func fifoParallelTest(t *testing.T, q Queue) {
	done := timeout(t)

	testRoutines := runtime.NumCPU()
	N := testItemCount / testRoutines
	wg := &sync.WaitGroup{}

	for k := 0; k < testRoutines; k += 1 {
		wg.Add(1)
		go func() {
			for i := 0; i < N; i += 1 {
				q.Enqueue(i)
			}

			for i := 0; i < N; i += 1 {
				_, ok := q.Dequeue()
				if !ok {
					t.Fatalf("Was unable to get value")
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	done <- 1
}

func lifoParallelTest(t *testing.T, q Queue) {
	done := timeout(t)

	testRoutines := runtime.NumCPU()
	N := testItemCount / testRoutines
	wg := &sync.WaitGroup{}
	for k := 0; k < testRoutines; k += 1 {
		wg.Add(1)
		go func() {
			for i := 0; i < N; i += 1 {
				q.Enqueue(i)
			}
			for i := 0; i < N; i += 1 {
				_, ok := q.Dequeue()
				if !ok {
					t.Fatalf("Was unable to get value")
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	done <- 1
}

func queueBench(b *testing.B, q Queue) {
	for i := 0; i < b.N; i += 1 {
		q.Enqueue(i)
	}
	for i := 0; i < b.N; i += 1 {
		q.Dequeue()
	}
}

func queueBenchParallel(b *testing.B, q Queue) {
	testRoutines := runtime.GOMAXPROCS(-1)
	wg := &sync.WaitGroup{}
	for k := 0; k < testRoutines; k += 1 {
		wg.Add(1)
		go func() {
			for i := 0; i < b.N; i += 1 {
				q.Enqueue(i)
			}
			for i := 0; i < b.N; i += 1 {
				q.Dequeue()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
