package goqueuestest

import (
	"testing"
)

func TestListChanFifo(t *testing.T) {
	fifoTest(t, NewListCFifo())
	fifoParallelTest(t, NewListCFifo())
}

func BenchmarkListChanFifo(b *testing.B) {
	queueBench(b, NewListCFifo())
}

func BenchmarkListChanFifoParallel(b *testing.B) {
	queueBenchParallel(b, NewListCFifo())
}

func TestListChanLifo(t *testing.T) {
	lifoTest(t, NewListCLifo())
	lifoParallelTest(t, NewListCLifo())
}

func BenchmarkListChanLifo(b *testing.B) {
	queueBench(b, NewListCLifo())
}

func BenchmarkListChanLifoParallel(b *testing.B) {
	queueBenchParallel(b, NewListCLifo())
}

func TestListMutexFifo(t *testing.T) {
	fifoTest(t, NewListFifo())
	fifoParallelTest(t, NewListFifo())
}

func BenchmarkListMutexFifo(b *testing.B) {
	queueBench(b, NewListFifo())
}

func BenchmarkListMutexFifoParallel(b *testing.B) {
	queueBenchParallel(b, NewListFifo())
}

func TestListMutexLifo(t *testing.T) {
	lifoTest(t, NewListLifo())
	lifoParallelTest(t, NewListLifo())
}

func BenchmarkListMutexLifo(b *testing.B) {
	queueBench(b, NewListLifo())
}

func BenchmarkListMutexLifoParallel(b *testing.B) {
	queueBenchParallel(b, NewListLifo())
}

func TestCFifo(t *testing.T) {
	fifoTest(t, NewChanFifo(testItemCount))
	fifoParallelTest(t, NewChanFifo(testItemCount))
}

func BenchmarkChanFifo(b *testing.B) {
	queueBench(b, NewChanFifo(b.N))
}

func BenchmarkChanFifoParallel(b *testing.B) {
	queueBenchParallel(b, NewChanFifo(b.N))
}

func TestZFifo(t *testing.T) {
	fifoTest(t, NewZFifo())
	fifoParallelTest(t, NewZFifo())
}

func BenchmarkLockfreeFifo(b *testing.B) {
	queueBench(b, NewZFifo())
}

func BenchmarkLockfreeFifoParallel(b *testing.B) {
	queueBenchParallel(b, NewZFifo())
}


func TestZLifo(t *testing.T) {
	lifoTest(t, NewZLifo())
	lifoParallelTest(t, NewZLifo())
}

func BenchmarkLockfreeLifo(b *testing.B) {
	queueBench(b, NewZLifo())
}

func BenchmarkLockfreeLifoParallel(b *testing.B) {
	queueBenchParallel(b, NewZLifo())
}


func TestZFifoFreechan(t *testing.T) {
	fifoTest(t, NewZFifoFreechan())
	fifoParallelTest(t, NewZFifoFreechan())
}

func BenchmarkLockfreeFifoFreechan(b *testing.B) {
	queueBench(b, NewZFifoFreechan())
}

func BenchmarkLockfreeFifoFreechanParallel(b *testing.B) {
	queueBenchParallel(b, NewZFifoFreechan())
}

func TestZFifoFreering(t *testing.T) {
	fifoTest(t, NewZFifoFreering())
	fifoParallelTest(t, NewZFifoFreering())
}

func BenchmarkLockfreeFifoFreering(b *testing.B) {
	queueBench(b, NewZFifoFreering())
}

func BenchmarkLockfreeFifoFreeringParallel(b *testing.B) {
	queueBenchParallel(b, NewZFifoFreering())
}
