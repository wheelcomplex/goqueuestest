package goqueuestest

import (
	"testing"
)

// SIMPLE TESTS

func TestListChanFifo(t *testing.T) {
	fifoTest(t, NewListCFifo())
	queuePTest(t, NewListCFifo())
	queueP2Test(t, NewListCFifo())
}

func TestListChanLifo(t *testing.T) {
	lifoTest(t, NewListCLifo())
	queuePTest(t, NewListCLifo())
	queueP2Test(t, NewListCLifo())
}

func TestListMutexFifo(t *testing.T) {
	fifoTest(t, NewListFifo())
	queuePTest(t, NewListFifo())
	queueP2Test(t, NewListFifo())
}

func TestListMutexLifo(t *testing.T) {
	lifoTest(t, NewListLifo())
	queuePTest(t, NewListLifo())
	queueP2Test(t, NewListLifo())
}

func TestZFifo(t *testing.T) {
	fifoTest(t, NewZFifo())
	queuePTest(t, NewZFifo())
	queueP2Test(t, NewZFifo())
}

func TestZLifo(t *testing.T) {
	lifoTest(t, NewZLifo())
	queuePTest(t, NewZLifo())
	queueP2Test(t, NewZLifo())
}

func TestZFifoFreechan(t *testing.T) {
	fifoTest(t, NewZFifoFreechan())
	queuePTest(t, NewZFifoFreechan())
	queueP2Test(t, NewZFifoFreechan())
}

func TestZFifoFreering(t *testing.T) {
	fifoTest(t, NewZFifoFreering())
	queuePTest(t, NewZFifoFreering())
	queueP2Test(t, NewZFifoFreering())
}

func TestRingLifo(t *testing.T) {
	lifoTest(t, NewRLifo())
	queuePTest(t, NewRLifo())
	queueP2Test(t, NewRLifo())
}

func TestRingFifo(t *testing.T) {
	fifoTest(t, NewRFifo())
	queuePTest(t, NewRFifo())
	queueP2Test(t, NewRFifo())
}

func TestSliceLifo(t *testing.T) {
	lifoTest(t, NewSLifo())
	queuePTest(t, NewSLifo())
	queueP2Test(t, NewSLifo())
}

func TestSliceFifo(t *testing.T) {
	fifoTest(t, NewSFifo())
	queuePTest(t, NewSFifo())
	queueP2Test(t, NewSFifo())
}

// SIMPLE BENCHMARK

func BenchmarkListChanFifo(b *testing.B) {
	queueBench(b, NewListCFifo())
}

func BenchmarkListChanLifo(b *testing.B) {
	queueBench(b, NewListCLifo())
}

func BenchmarkListMutexFifo(b *testing.B) {
	queueBench(b, NewListFifo())
}

func BenchmarkListMutexLifo(b *testing.B) {
	queueBench(b, NewListLifo())
}


func BenchmarkLockfreeFifo(b *testing.B) {
	queueBench(b, NewZFifo())
}

func BenchmarkLockfreeLifo(b *testing.B) {
	queueBench(b, NewZLifo())
}

func BenchmarkLockfreeFifoFreechan(b *testing.B) {
	queueBench(b, NewZFifoFreechan())
}

func BenchmarkLockfreeFifoFreering(b *testing.B) {
	queueBench(b, NewZFifoFreering())
}

func BenchmarkRingLifo(b *testing.B) {
	queueBench(b, NewRLifo())
}

func BenchmarkRingFifo(b *testing.B) {
	queueBench(b, NewRFifo())
}

func BenchmarkSliceLifo(b *testing.B) {
	queueBench(b, NewSLifo())
}

func BenchmarkSliceFifo(b *testing.B) {
	queueBench(b, NewSFifo())
}

// PARALLEL BENCHMARK

func BenchmarkListChanLifoP(b *testing.B) {
	queueBenchP(b, NewListCLifo())
}

func BenchmarkListMutexFifoP(b *testing.B) {
	queueBenchP(b, NewListFifo())
}

func BenchmarkListMutexLifoP(b *testing.B) {
	queueBenchP(b, NewListLifo())
}

func BenchmarkLockfreeFifoP(b *testing.B) {
	queueBenchP(b, NewZFifo())
}

func BenchmarkLockfreeLifoP(b *testing.B) {
	queueBenchP(b, NewZLifo())
}

func BenchmarkLockfreeFifoFreechanP(b *testing.B) {
	queueBenchP(b, NewZFifoFreechan())
}

func BenchmarkLockfreeFifoFreeringP(b *testing.B) {
	queueBenchP(b, NewZFifoFreering())
}

func BenchmarkRingLifoP(b *testing.B) {
	queueBenchP(b, NewRLifo())
}

func BenchmarkRingFifoP(b *testing.B) {
	queueBenchP(b, NewRFifo())
}

func BenchmarkSliceLifoP(b *testing.B) {
	queueBenchP(b, NewSLifo())
}

func BenchmarkSliceFifoP(b *testing.B) {
	queueBenchP(b, NewSFifo())
}

// PARALLEL BENCHMARK 2

func BenchmarkListChanFifoP2(b *testing.B) {
	queueBenchP2(b, NewListCFifo())
}

func BenchmarkListChanLifoP2(b *testing.B) {
	queueBenchP2(b, NewListCLifo())
}

func BenchmarkListMutexFifoP2(b *testing.B) {
	queueBenchP2(b, NewListFifo())
}

func BenchmarkListMutexLifoP2(b *testing.B) {
	queueBenchP2(b, NewListLifo())
}

func BenchmarkChanFifoP2(b *testing.B) {
	queueBenchP2(b, NewChanFifo(b.N))
}

func BenchmarkLockfreeFifoP2(b *testing.B) {
	queueBenchP2(b, NewZFifo())
}

func BenchmarkLockfreeLifoP2(b *testing.B) {
	queueBenchP2(b, NewZLifo())
}

func BrokenBenchmarkLockfreeFifoFreechanP2(b *testing.B) {
	queueBenchP2(b, NewZFifoFreechan())
}

func BenchmarkLockfreeFifoFreeringP2(b *testing.B) {
	queueBenchP2(b, NewZFifoFreering())
}

func BenchmarkRingLifoP2(b *testing.B) {
	queueBenchP2(b, NewRLifo())
}

func BenchmarkRingFifoP2(b *testing.B) {
	queueBenchP2(b, NewRFifo())
}

func BenchmarkSliceLifoP2(b *testing.B) {
	queueBenchP2(b, NewSLifo())
}

func BenchmarkSliceFifoP2(b *testing.B) {
	queueBenchP2(b, NewSFifo())
}