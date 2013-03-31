package goqueuestest
// DO NOT EDIT, GENERATED CODE
import "testing"


func TestLmLifo(t *testing.T) {
	lifoTest(t, NewListLifo())
	queuePTest(t, NewListLifo())
	queueP2Test(t, NewListLifo())
}

func BenchmarkLmLifo(b *testing.B) {
	queueBench(b, NewListLifo())
}

func BenchmarkLmLifoP(b *testing.B) {
	queueBenchP(b, NewListLifo())
}

func BenchmarkLmLifoP2(b *testing.B) {
	queueBenchP2(b, NewListLifo())
}

func TestRmLifo(t *testing.T) {
	lifoTest(t, NewRLifo())
	queuePTest(t, NewRLifo())
	queueP2Test(t, NewRLifo())
}

func BenchmarkRmLifo(b *testing.B) {
	queueBench(b, NewRLifo())
}

func BenchmarkRmLifoP(b *testing.B) {
	queueBenchP(b, NewRLifo())
}

func BenchmarkRmLifoP2(b *testing.B) {
	queueBenchP2(b, NewRLifo())
}

func TestRmFifo(t *testing.T) {
	fifoTest(t, NewRFifo())
	queuePTest(t, NewRFifo())
	queueP2Test(t, NewRFifo())
}

func BenchmarkRmFifo(b *testing.B) {
	queueBench(b, NewRFifo())
}

func BenchmarkRmFifoP(b *testing.B) {
	queueBenchP(b, NewRFifo())
}

func BenchmarkRmFifoP2(b *testing.B) {
	queueBenchP2(b, NewRFifo())
}

func TestSmLifo(t *testing.T) {
	lifoTest(t, NewSLifo())
	queuePTest(t, NewSLifo())
	queueP2Test(t, NewSLifo())
}

func BenchmarkSmLifo(b *testing.B) {
	queueBench(b, NewSLifo())
}

func BenchmarkSmLifoP(b *testing.B) {
	queueBenchP(b, NewSLifo())
}

func BenchmarkSmLifoP2(b *testing.B) {
	queueBenchP2(b, NewSLifo())
}

func TestCFifo(t *testing.T) {
	fifoTest(t, NewChanFifo(testItemCount))
	queuePTest(t, NewChanFifo(testItemCount))
	queueP2Test(t, NewChanFifo(testItemCount))
}

func BenchmarkCFifoS(b *testing.B) {
	queueBench(b, NewChanFifo(b.N))
}

func BenchmarkCFifoP(b *testing.B) {
	queueBenchP(b, NewChanFifo(b.N))
}

func BenchmarkCFifoP2(b *testing.B) {
	queueBenchP2(b, NewChanFifo(b.N))
}

func TestLcLifo(t *testing.T) {
	lifoTest(t, NewListCLifo())
	queuePTest(t, NewListCLifo())
	queueP2Test(t, NewListCLifo())
}

func BenchmarkLcLifo(b *testing.B) {
	queueBench(b, NewListCLifo())
}

func BenchmarkLcLifoP(b *testing.B) {
	queueBenchP(b, NewListCLifo())
}

func BenchmarkLcLifoP2(b *testing.B) {
	queueBenchP2(b, NewListCLifo())
}

func TestLcFifo(t *testing.T) {
	fifoTest(t, NewListCFifo())
	queuePTest(t, NewListCFifo())
	queueP2Test(t, NewListCFifo())
}

func BenchmarkLcFifo(b *testing.B) {
	queueBench(b, NewListCFifo())
}

func BenchmarkLcFifoP(b *testing.B) {
	queueBenchP(b, NewListCFifo())
}

func BenchmarkLcFifoP2(b *testing.B) {
	queueBenchP2(b, NewListCFifo())
}

func TestLmFifo(t *testing.T) {
	fifoTest(t, NewListFifo())
	queuePTest(t, NewListFifo())
	queueP2Test(t, NewListFifo())
}

func BenchmarkLmFifo(b *testing.B) {
	queueBench(b, NewListFifo())
}

func BenchmarkLmFifoP(b *testing.B) {
	queueBenchP(b, NewListFifo())
}

func BenchmarkLmFifoP2(b *testing.B) {
	queueBenchP2(b, NewListFifo())
}

func TestZLifo(t *testing.T) {
	lifoTest(t, NewZLifo())
	queuePTest(t, NewZLifo())
	queueP2Test(t, NewZLifo())
}

func BenchmarkZLifo(b *testing.B) {
	queueBench(b, NewZLifo())
}

func BenchmarkZLifoP(b *testing.B) {
	queueBenchP(b, NewZLifo())
}

func BenchmarkZLifoP2(b *testing.B) {
	queueBenchP2(b, NewZLifo())
}

func TestZFifo(t *testing.T) {
	fifoTest(t, NewZFifo())
	queuePTest(t, NewZFifo())
	queueP2Test(t, NewZFifo())
}

func BenchmarkZFifo(b *testing.B) {
	queueBench(b, NewZFifo())
}

func BenchmarkZFifoP(b *testing.B) {
	queueBenchP(b, NewZFifo())
}

func BenchmarkZFifoP2(b *testing.B) {
	queueBenchP2(b, NewZFifo())
}

func TestZcFifo(t *testing.T) {
	fifoTest(t, NewZFifoFreechan())
	queuePTest(t, NewZFifoFreechan())
	queueP2Test(t, NewZFifoFreechan())
}

func BenchmarkZcFifo(b *testing.B) {
	queueBench(b, NewZFifoFreechan())
}

func BenchmarkZcFifoP(b *testing.B) {
	queueBenchP(b, NewZFifoFreechan())
}

func BenchmarkZcFifoP2(b *testing.B) {
	queueBenchP2(b, NewZFifoFreechan())
}

func TestZrFifo(t *testing.T) {
	fifoTest(t, NewZFifoFreering())
	queuePTest(t, NewZFifoFreering())
	queueP2Test(t, NewZFifoFreering())
}

func BenchmarkZrFifo(b *testing.B) {
	queueBench(b, NewZFifoFreering())
}

func BenchmarkZrFifoP(b *testing.B) {
	queueBenchP(b, NewZFifoFreering())
}

func BenchmarkZrFifoP2(b *testing.B) {
	queueBenchP2(b, NewZFifoFreering())
}

func TestSmFifo(t *testing.T) {
	fifoTest(t, NewSFifo())
	queuePTest(t, NewSFifo())
	queueP2Test(t, NewSFifo())
}

func BenchmarkSmFifo(b *testing.B) {
	queueBench(b, NewSFifo())
}

func BenchmarkSmFifoP(b *testing.B) {
	queueBenchP(b, NewSFifo())
}

func BenchmarkSmFifoP2(b *testing.B) {
	queueBenchP2(b, NewSFifo())
}
