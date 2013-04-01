package goqueuestest
// DO NOT EDIT, GENERATED CODE
import "testing"


// CFifo

func TestCFifo(t *testing.T) {
	fifoTest(t, NewChanFifo(benchIterations))
	queuePTest(t, NewChanFifo(benchIterations))
	queueP2Test(t, NewChanFifo(benchIterations))
}
func BenchmarkCFifo_ANRNx1(b *testing.B) { runParallelBench(b, 1, NewChanFifo(benchIterations), benchANRN)}
func BenchmarkCFifo_ANRNx3(b *testing.B) { runParallelBench(b, 3, NewChanFifo(benchIterations), benchANRN)}
func BenchmarkCFifo_ANRNx5(b *testing.B) { runParallelBench(b, 5, NewChanFifo(benchIterations), benchANRN)}
func BenchmarkCFifo_ANRNx10(b *testing.B) { runParallelBench(b, 10, NewChanFifo(benchIterations), benchANRN)}
func BenchmarkCFifo_ANRNx50(b *testing.B) { runParallelBench(b, 50, NewChanFifo(benchIterations), benchANRN)}
func BenchmarkCFifo_ANRNx100(b *testing.B) { runParallelBench(b, 100, NewChanFifo(benchIterations), benchANRN)}
func BenchmarkCFifo_A1R1x1(b *testing.B) { runParallelBench(b, 1, NewChanFifo(benchIterations), benchA1R1)}
func BenchmarkCFifo_A1R1x3(b *testing.B) { runParallelBench(b, 3, NewChanFifo(benchIterations), benchA1R1)}
func BenchmarkCFifo_A1R1x5(b *testing.B) { runParallelBench(b, 5, NewChanFifo(benchIterations), benchA1R1)}
func BenchmarkCFifo_A1R1x10(b *testing.B) { runParallelBench(b, 10, NewChanFifo(benchIterations), benchA1R1)}
func BenchmarkCFifo_A1R1x50(b *testing.B) { runParallelBench(b, 50, NewChanFifo(benchIterations), benchA1R1)}
func BenchmarkCFifo_A1R1x100(b *testing.B) { runParallelBench(b, 100, NewChanFifo(benchIterations), benchA1R1)}
func BenchmarkCFifo_A1R2x1(b *testing.B) { runParallelBench(b, 1, NewChanFifo(benchIterations), benchA1R2)}
func BenchmarkCFifo_A1R2x3(b *testing.B) { runParallelBench(b, 3, NewChanFifo(benchIterations), benchA1R2)}
func BenchmarkCFifo_A1R2x5(b *testing.B) { runParallelBench(b, 5, NewChanFifo(benchIterations), benchA1R2)}
func BenchmarkCFifo_A1R2x10(b *testing.B) { runParallelBench(b, 10, NewChanFifo(benchIterations), benchA1R2)}
func BenchmarkCFifo_A1R2x50(b *testing.B) { runParallelBench(b, 50, NewChanFifo(benchIterations), benchA1R2)}
func BenchmarkCFifo_A1R2x100(b *testing.B) { runParallelBench(b, 100, NewChanFifo(benchIterations), benchA1R2)}
func BenchmarkCFifo_A2R1x1(b *testing.B) { runParallelBench(b, 1, NewChanFifo(benchIterations), benchA2R1)}
func BenchmarkCFifo_A2R1x3(b *testing.B) { runParallelBench(b, 3, NewChanFifo(benchIterations), benchA2R1)}
func BenchmarkCFifo_A2R1x5(b *testing.B) { runParallelBench(b, 5, NewChanFifo(benchIterations), benchA2R1)}
func BenchmarkCFifo_A2R1x10(b *testing.B) { runParallelBench(b, 10, NewChanFifo(benchIterations), benchA2R1)}
func BenchmarkCFifo_A2R1x50(b *testing.B) { runParallelBench(b, 50, NewChanFifo(benchIterations), benchA2R1)}
func BenchmarkCFifo_A2R1x100(b *testing.B) { runParallelBench(b, 100, NewChanFifo(benchIterations), benchA2R1)}

// ZLifo

func TestZLifo(t *testing.T) {
	lifoTest(t, NewZLifo())
	queuePTest(t, NewZLifo())
	queueP2Test(t, NewZLifo())
}
func BenchmarkZLifo_ANRNx1(b *testing.B) { runParallelBench(b, 1, NewZLifo(), benchANRN)}
func BenchmarkZLifo_ANRNx3(b *testing.B) { runParallelBench(b, 3, NewZLifo(), benchANRN)}
func BenchmarkZLifo_ANRNx5(b *testing.B) { runParallelBench(b, 5, NewZLifo(), benchANRN)}
func BenchmarkZLifo_ANRNx10(b *testing.B) { runParallelBench(b, 10, NewZLifo(), benchANRN)}
func BenchmarkZLifo_ANRNx50(b *testing.B) { runParallelBench(b, 50, NewZLifo(), benchANRN)}
func BenchmarkZLifo_ANRNx100(b *testing.B) { runParallelBench(b, 100, NewZLifo(), benchANRN)}
func BenchmarkZLifo_A1R1x1(b *testing.B) { runParallelBench(b, 1, NewZLifo(), benchA1R1)}
func BenchmarkZLifo_A1R1x3(b *testing.B) { runParallelBench(b, 3, NewZLifo(), benchA1R1)}
func BenchmarkZLifo_A1R1x5(b *testing.B) { runParallelBench(b, 5, NewZLifo(), benchA1R1)}
func BenchmarkZLifo_A1R1x10(b *testing.B) { runParallelBench(b, 10, NewZLifo(), benchA1R1)}
func BenchmarkZLifo_A1R1x50(b *testing.B) { runParallelBench(b, 50, NewZLifo(), benchA1R1)}
func BenchmarkZLifo_A1R1x100(b *testing.B) { runParallelBench(b, 100, NewZLifo(), benchA1R1)}
func BenchmarkZLifo_A1R2x1(b *testing.B) { runParallelBench(b, 1, NewZLifo(), benchA1R2)}
func BenchmarkZLifo_A1R2x3(b *testing.B) { runParallelBench(b, 3, NewZLifo(), benchA1R2)}
func BenchmarkZLifo_A1R2x5(b *testing.B) { runParallelBench(b, 5, NewZLifo(), benchA1R2)}
func BenchmarkZLifo_A1R2x10(b *testing.B) { runParallelBench(b, 10, NewZLifo(), benchA1R2)}
func BenchmarkZLifo_A1R2x50(b *testing.B) { runParallelBench(b, 50, NewZLifo(), benchA1R2)}
func BenchmarkZLifo_A1R2x100(b *testing.B) { runParallelBench(b, 100, NewZLifo(), benchA1R2)}
func BenchmarkZLifo_A2R1x1(b *testing.B) { runParallelBench(b, 1, NewZLifo(), benchA2R1)}
func BenchmarkZLifo_A2R1x3(b *testing.B) { runParallelBench(b, 3, NewZLifo(), benchA2R1)}
func BenchmarkZLifo_A2R1x5(b *testing.B) { runParallelBench(b, 5, NewZLifo(), benchA2R1)}
func BenchmarkZLifo_A2R1x10(b *testing.B) { runParallelBench(b, 10, NewZLifo(), benchA2R1)}
func BenchmarkZLifo_A2R1x50(b *testing.B) { runParallelBench(b, 50, NewZLifo(), benchA2R1)}
func BenchmarkZLifo_A2R1x100(b *testing.B) { runParallelBench(b, 100, NewZLifo(), benchA2R1)}

// ZFifo

func TestZFifo(t *testing.T) {
	fifoTest(t, NewZFifo())
	queuePTest(t, NewZFifo())
	queueP2Test(t, NewZFifo())
}
func BenchmarkZFifo_ANRNx1(b *testing.B) { runParallelBench(b, 1, NewZFifo(), benchANRN)}
func BenchmarkZFifo_ANRNx3(b *testing.B) { runParallelBench(b, 3, NewZFifo(), benchANRN)}
func BenchmarkZFifo_ANRNx5(b *testing.B) { runParallelBench(b, 5, NewZFifo(), benchANRN)}
func BenchmarkZFifo_ANRNx10(b *testing.B) { runParallelBench(b, 10, NewZFifo(), benchANRN)}
func BenchmarkZFifo_ANRNx50(b *testing.B) { runParallelBench(b, 50, NewZFifo(), benchANRN)}
func BenchmarkZFifo_ANRNx100(b *testing.B) { runParallelBench(b, 100, NewZFifo(), benchANRN)}
func BenchmarkZFifo_A1R1x1(b *testing.B) { runParallelBench(b, 1, NewZFifo(), benchA1R1)}
func BenchmarkZFifo_A1R1x3(b *testing.B) { runParallelBench(b, 3, NewZFifo(), benchA1R1)}
func BenchmarkZFifo_A1R1x5(b *testing.B) { runParallelBench(b, 5, NewZFifo(), benchA1R1)}
func BenchmarkZFifo_A1R1x10(b *testing.B) { runParallelBench(b, 10, NewZFifo(), benchA1R1)}
func BenchmarkZFifo_A1R1x50(b *testing.B) { runParallelBench(b, 50, NewZFifo(), benchA1R1)}
func BenchmarkZFifo_A1R1x100(b *testing.B) { runParallelBench(b, 100, NewZFifo(), benchA1R1)}
func BenchmarkZFifo_A1R2x1(b *testing.B) { runParallelBench(b, 1, NewZFifo(), benchA1R2)}
func BenchmarkZFifo_A1R2x3(b *testing.B) { runParallelBench(b, 3, NewZFifo(), benchA1R2)}
func BenchmarkZFifo_A1R2x5(b *testing.B) { runParallelBench(b, 5, NewZFifo(), benchA1R2)}
func BenchmarkZFifo_A1R2x10(b *testing.B) { runParallelBench(b, 10, NewZFifo(), benchA1R2)}
func BenchmarkZFifo_A1R2x50(b *testing.B) { runParallelBench(b, 50, NewZFifo(), benchA1R2)}
func BenchmarkZFifo_A1R2x100(b *testing.B) { runParallelBench(b, 100, NewZFifo(), benchA1R2)}
func BenchmarkZFifo_A2R1x1(b *testing.B) { runParallelBench(b, 1, NewZFifo(), benchA2R1)}
func BenchmarkZFifo_A2R1x3(b *testing.B) { runParallelBench(b, 3, NewZFifo(), benchA2R1)}
func BenchmarkZFifo_A2R1x5(b *testing.B) { runParallelBench(b, 5, NewZFifo(), benchA2R1)}
func BenchmarkZFifo_A2R1x10(b *testing.B) { runParallelBench(b, 10, NewZFifo(), benchA2R1)}
func BenchmarkZFifo_A2R1x50(b *testing.B) { runParallelBench(b, 50, NewZFifo(), benchA2R1)}
func BenchmarkZFifo_A2R1x100(b *testing.B) { runParallelBench(b, 100, NewZFifo(), benchA2R1)}

// ScLifo

func TestScLifo(t *testing.T) {
	lifoTest(t, NewScLifo())
	queuePTest(t, NewScLifo())
	queueP2Test(t, NewScLifo())
}
func BenchmarkScLifo_ANRNx1(b *testing.B) { runParallelBench(b, 1, NewScLifo(), benchANRN)}
func BenchmarkScLifo_ANRNx3(b *testing.B) { runParallelBench(b, 3, NewScLifo(), benchANRN)}
func BenchmarkScLifo_ANRNx5(b *testing.B) { runParallelBench(b, 5, NewScLifo(), benchANRN)}
func BenchmarkScLifo_ANRNx10(b *testing.B) { runParallelBench(b, 10, NewScLifo(), benchANRN)}
func BenchmarkScLifo_ANRNx50(b *testing.B) { runParallelBench(b, 50, NewScLifo(), benchANRN)}
func BenchmarkScLifo_ANRNx100(b *testing.B) { runParallelBench(b, 100, NewScLifo(), benchANRN)}
func BenchmarkScLifo_A1R1x1(b *testing.B) { runParallelBench(b, 1, NewScLifo(), benchA1R1)}
func BenchmarkScLifo_A1R1x3(b *testing.B) { runParallelBench(b, 3, NewScLifo(), benchA1R1)}
func BenchmarkScLifo_A1R1x5(b *testing.B) { runParallelBench(b, 5, NewScLifo(), benchA1R1)}
func BenchmarkScLifo_A1R1x10(b *testing.B) { runParallelBench(b, 10, NewScLifo(), benchA1R1)}
func BenchmarkScLifo_A1R1x50(b *testing.B) { runParallelBench(b, 50, NewScLifo(), benchA1R1)}
func BenchmarkScLifo_A1R1x100(b *testing.B) { runParallelBench(b, 100, NewScLifo(), benchA1R1)}
func BenchmarkScLifo_A1R2x1(b *testing.B) { runParallelBench(b, 1, NewScLifo(), benchA1R2)}
func BenchmarkScLifo_A1R2x3(b *testing.B) { runParallelBench(b, 3, NewScLifo(), benchA1R2)}
func BenchmarkScLifo_A1R2x5(b *testing.B) { runParallelBench(b, 5, NewScLifo(), benchA1R2)}
func BenchmarkScLifo_A1R2x10(b *testing.B) { runParallelBench(b, 10, NewScLifo(), benchA1R2)}
func BenchmarkScLifo_A1R2x50(b *testing.B) { runParallelBench(b, 50, NewScLifo(), benchA1R2)}
func BenchmarkScLifo_A1R2x100(b *testing.B) { runParallelBench(b, 100, NewScLifo(), benchA1R2)}
func BenchmarkScLifo_A2R1x1(b *testing.B) { runParallelBench(b, 1, NewScLifo(), benchA2R1)}
func BenchmarkScLifo_A2R1x3(b *testing.B) { runParallelBench(b, 3, NewScLifo(), benchA2R1)}
func BenchmarkScLifo_A2R1x5(b *testing.B) { runParallelBench(b, 5, NewScLifo(), benchA2R1)}
func BenchmarkScLifo_A2R1x10(b *testing.B) { runParallelBench(b, 10, NewScLifo(), benchA2R1)}
func BenchmarkScLifo_A2R1x50(b *testing.B) { runParallelBench(b, 50, NewScLifo(), benchA2R1)}
func BenchmarkScLifo_A2R1x100(b *testing.B) { runParallelBench(b, 100, NewScLifo(), benchA2R1)}

// ScFifo

func TestScFifo(t *testing.T) {
	fifoTest(t, NewScFifo())
	queuePTest(t, NewScFifo())
	queueP2Test(t, NewScFifo())
}
func BenchmarkScFifo_ANRNx1(b *testing.B) { runParallelBench(b, 1, NewScFifo(), benchANRN)}
func BenchmarkScFifo_ANRNx3(b *testing.B) { runParallelBench(b, 3, NewScFifo(), benchANRN)}
func BenchmarkScFifo_ANRNx5(b *testing.B) { runParallelBench(b, 5, NewScFifo(), benchANRN)}
func BenchmarkScFifo_ANRNx10(b *testing.B) { runParallelBench(b, 10, NewScFifo(), benchANRN)}
func BenchmarkScFifo_ANRNx50(b *testing.B) { runParallelBench(b, 50, NewScFifo(), benchANRN)}
func BenchmarkScFifo_ANRNx100(b *testing.B) { runParallelBench(b, 100, NewScFifo(), benchANRN)}
func BenchmarkScFifo_A1R1x1(b *testing.B) { runParallelBench(b, 1, NewScFifo(), benchA1R1)}
func BenchmarkScFifo_A1R1x3(b *testing.B) { runParallelBench(b, 3, NewScFifo(), benchA1R1)}
func BenchmarkScFifo_A1R1x5(b *testing.B) { runParallelBench(b, 5, NewScFifo(), benchA1R1)}
func BenchmarkScFifo_A1R1x10(b *testing.B) { runParallelBench(b, 10, NewScFifo(), benchA1R1)}
func BenchmarkScFifo_A1R1x50(b *testing.B) { runParallelBench(b, 50, NewScFifo(), benchA1R1)}
func BenchmarkScFifo_A1R1x100(b *testing.B) { runParallelBench(b, 100, NewScFifo(), benchA1R1)}
func BenchmarkScFifo_A1R2x1(b *testing.B) { runParallelBench(b, 1, NewScFifo(), benchA1R2)}
func BenchmarkScFifo_A1R2x3(b *testing.B) { runParallelBench(b, 3, NewScFifo(), benchA1R2)}
func BenchmarkScFifo_A1R2x5(b *testing.B) { runParallelBench(b, 5, NewScFifo(), benchA1R2)}
func BenchmarkScFifo_A1R2x10(b *testing.B) { runParallelBench(b, 10, NewScFifo(), benchA1R2)}
func BenchmarkScFifo_A1R2x50(b *testing.B) { runParallelBench(b, 50, NewScFifo(), benchA1R2)}
func BenchmarkScFifo_A1R2x100(b *testing.B) { runParallelBench(b, 100, NewScFifo(), benchA1R2)}
func BenchmarkScFifo_A2R1x1(b *testing.B) { runParallelBench(b, 1, NewScFifo(), benchA2R1)}
func BenchmarkScFifo_A2R1x3(b *testing.B) { runParallelBench(b, 3, NewScFifo(), benchA2R1)}
func BenchmarkScFifo_A2R1x5(b *testing.B) { runParallelBench(b, 5, NewScFifo(), benchA2R1)}
func BenchmarkScFifo_A2R1x10(b *testing.B) { runParallelBench(b, 10, NewScFifo(), benchA2R1)}
func BenchmarkScFifo_A2R1x50(b *testing.B) { runParallelBench(b, 50, NewScFifo(), benchA2R1)}
func BenchmarkScFifo_A2R1x100(b *testing.B) { runParallelBench(b, 100, NewScFifo(), benchA2R1)}

// SmLifo

func TestSmLifo(t *testing.T) {
	lifoTest(t, NewSmLifo())
	queuePTest(t, NewSmLifo())
	queueP2Test(t, NewSmLifo())
}
func BenchmarkSmLifo_ANRNx1(b *testing.B) { runParallelBench(b, 1, NewSmLifo(), benchANRN)}
func BenchmarkSmLifo_ANRNx3(b *testing.B) { runParallelBench(b, 3, NewSmLifo(), benchANRN)}
func BenchmarkSmLifo_ANRNx5(b *testing.B) { runParallelBench(b, 5, NewSmLifo(), benchANRN)}
func BenchmarkSmLifo_ANRNx10(b *testing.B) { runParallelBench(b, 10, NewSmLifo(), benchANRN)}
func BenchmarkSmLifo_ANRNx50(b *testing.B) { runParallelBench(b, 50, NewSmLifo(), benchANRN)}
func BenchmarkSmLifo_ANRNx100(b *testing.B) { runParallelBench(b, 100, NewSmLifo(), benchANRN)}
func BenchmarkSmLifo_A1R1x1(b *testing.B) { runParallelBench(b, 1, NewSmLifo(), benchA1R1)}
func BenchmarkSmLifo_A1R1x3(b *testing.B) { runParallelBench(b, 3, NewSmLifo(), benchA1R1)}
func BenchmarkSmLifo_A1R1x5(b *testing.B) { runParallelBench(b, 5, NewSmLifo(), benchA1R1)}
func BenchmarkSmLifo_A1R1x10(b *testing.B) { runParallelBench(b, 10, NewSmLifo(), benchA1R1)}
func BenchmarkSmLifo_A1R1x50(b *testing.B) { runParallelBench(b, 50, NewSmLifo(), benchA1R1)}
func BenchmarkSmLifo_A1R1x100(b *testing.B) { runParallelBench(b, 100, NewSmLifo(), benchA1R1)}
func BenchmarkSmLifo_A1R2x1(b *testing.B) { runParallelBench(b, 1, NewSmLifo(), benchA1R2)}
func BenchmarkSmLifo_A1R2x3(b *testing.B) { runParallelBench(b, 3, NewSmLifo(), benchA1R2)}
func BenchmarkSmLifo_A1R2x5(b *testing.B) { runParallelBench(b, 5, NewSmLifo(), benchA1R2)}
func BenchmarkSmLifo_A1R2x10(b *testing.B) { runParallelBench(b, 10, NewSmLifo(), benchA1R2)}
func BenchmarkSmLifo_A1R2x50(b *testing.B) { runParallelBench(b, 50, NewSmLifo(), benchA1R2)}
func BenchmarkSmLifo_A1R2x100(b *testing.B) { runParallelBench(b, 100, NewSmLifo(), benchA1R2)}
func BenchmarkSmLifo_A2R1x1(b *testing.B) { runParallelBench(b, 1, NewSmLifo(), benchA2R1)}
func BenchmarkSmLifo_A2R1x3(b *testing.B) { runParallelBench(b, 3, NewSmLifo(), benchA2R1)}
func BenchmarkSmLifo_A2R1x5(b *testing.B) { runParallelBench(b, 5, NewSmLifo(), benchA2R1)}
func BenchmarkSmLifo_A2R1x10(b *testing.B) { runParallelBench(b, 10, NewSmLifo(), benchA2R1)}
func BenchmarkSmLifo_A2R1x50(b *testing.B) { runParallelBench(b, 50, NewSmLifo(), benchA2R1)}
func BenchmarkSmLifo_A2R1x100(b *testing.B) { runParallelBench(b, 100, NewSmLifo(), benchA2R1)}

// SmFifo

func TestSmFifo(t *testing.T) {
	fifoTest(t, NewSmFifo())
	queuePTest(t, NewSmFifo())
	queueP2Test(t, NewSmFifo())
}
func BenchmarkSmFifo_ANRNx1(b *testing.B) { runParallelBench(b, 1, NewSmFifo(), benchANRN)}
func BenchmarkSmFifo_ANRNx3(b *testing.B) { runParallelBench(b, 3, NewSmFifo(), benchANRN)}
func BenchmarkSmFifo_ANRNx5(b *testing.B) { runParallelBench(b, 5, NewSmFifo(), benchANRN)}
func BenchmarkSmFifo_ANRNx10(b *testing.B) { runParallelBench(b, 10, NewSmFifo(), benchANRN)}
func BenchmarkSmFifo_ANRNx50(b *testing.B) { runParallelBench(b, 50, NewSmFifo(), benchANRN)}
func BenchmarkSmFifo_ANRNx100(b *testing.B) { runParallelBench(b, 100, NewSmFifo(), benchANRN)}
func BenchmarkSmFifo_A1R1x1(b *testing.B) { runParallelBench(b, 1, NewSmFifo(), benchA1R1)}
func BenchmarkSmFifo_A1R1x3(b *testing.B) { runParallelBench(b, 3, NewSmFifo(), benchA1R1)}
func BenchmarkSmFifo_A1R1x5(b *testing.B) { runParallelBench(b, 5, NewSmFifo(), benchA1R1)}
func BenchmarkSmFifo_A1R1x10(b *testing.B) { runParallelBench(b, 10, NewSmFifo(), benchA1R1)}
func BenchmarkSmFifo_A1R1x50(b *testing.B) { runParallelBench(b, 50, NewSmFifo(), benchA1R1)}
func BenchmarkSmFifo_A1R1x100(b *testing.B) { runParallelBench(b, 100, NewSmFifo(), benchA1R1)}
func BenchmarkSmFifo_A1R2x1(b *testing.B) { runParallelBench(b, 1, NewSmFifo(), benchA1R2)}
func BenchmarkSmFifo_A1R2x3(b *testing.B) { runParallelBench(b, 3, NewSmFifo(), benchA1R2)}
func BenchmarkSmFifo_A1R2x5(b *testing.B) { runParallelBench(b, 5, NewSmFifo(), benchA1R2)}
func BenchmarkSmFifo_A1R2x10(b *testing.B) { runParallelBench(b, 10, NewSmFifo(), benchA1R2)}
func BenchmarkSmFifo_A1R2x50(b *testing.B) { runParallelBench(b, 50, NewSmFifo(), benchA1R2)}
func BenchmarkSmFifo_A1R2x100(b *testing.B) { runParallelBench(b, 100, NewSmFifo(), benchA1R2)}
func BenchmarkSmFifo_A2R1x1(b *testing.B) { runParallelBench(b, 1, NewSmFifo(), benchA2R1)}
func BenchmarkSmFifo_A2R1x3(b *testing.B) { runParallelBench(b, 3, NewSmFifo(), benchA2R1)}
func BenchmarkSmFifo_A2R1x5(b *testing.B) { runParallelBench(b, 5, NewSmFifo(), benchA2R1)}
func BenchmarkSmFifo_A2R1x10(b *testing.B) { runParallelBench(b, 10, NewSmFifo(), benchA2R1)}
func BenchmarkSmFifo_A2R1x50(b *testing.B) { runParallelBench(b, 50, NewSmFifo(), benchA2R1)}
func BenchmarkSmFifo_A2R1x100(b *testing.B) { runParallelBench(b, 100, NewSmFifo(), benchA2R1)}

// RcLifo

func TestRcLifo(t *testing.T) {
	lifoTest(t, NewRcLifo())
	queuePTest(t, NewRcLifo())
	queueP2Test(t, NewRcLifo())
}
func BenchmarkRcLifo_ANRNx1(b *testing.B) { runParallelBench(b, 1, NewRcLifo(), benchANRN)}
func BenchmarkRcLifo_ANRNx3(b *testing.B) { runParallelBench(b, 3, NewRcLifo(), benchANRN)}
func BenchmarkRcLifo_ANRNx5(b *testing.B) { runParallelBench(b, 5, NewRcLifo(), benchANRN)}
func BenchmarkRcLifo_ANRNx10(b *testing.B) { runParallelBench(b, 10, NewRcLifo(), benchANRN)}
func BenchmarkRcLifo_ANRNx50(b *testing.B) { runParallelBench(b, 50, NewRcLifo(), benchANRN)}
func BenchmarkRcLifo_ANRNx100(b *testing.B) { runParallelBench(b, 100, NewRcLifo(), benchANRN)}
func BenchmarkRcLifo_A1R1x1(b *testing.B) { runParallelBench(b, 1, NewRcLifo(), benchA1R1)}
func BenchmarkRcLifo_A1R1x3(b *testing.B) { runParallelBench(b, 3, NewRcLifo(), benchA1R1)}
func BenchmarkRcLifo_A1R1x5(b *testing.B) { runParallelBench(b, 5, NewRcLifo(), benchA1R1)}
func BenchmarkRcLifo_A1R1x10(b *testing.B) { runParallelBench(b, 10, NewRcLifo(), benchA1R1)}
func BenchmarkRcLifo_A1R1x50(b *testing.B) { runParallelBench(b, 50, NewRcLifo(), benchA1R1)}
func BenchmarkRcLifo_A1R1x100(b *testing.B) { runParallelBench(b, 100, NewRcLifo(), benchA1R1)}
func BenchmarkRcLifo_A1R2x1(b *testing.B) { runParallelBench(b, 1, NewRcLifo(), benchA1R2)}
func BenchmarkRcLifo_A1R2x3(b *testing.B) { runParallelBench(b, 3, NewRcLifo(), benchA1R2)}
func BenchmarkRcLifo_A1R2x5(b *testing.B) { runParallelBench(b, 5, NewRcLifo(), benchA1R2)}
func BenchmarkRcLifo_A1R2x10(b *testing.B) { runParallelBench(b, 10, NewRcLifo(), benchA1R2)}
func BenchmarkRcLifo_A1R2x50(b *testing.B) { runParallelBench(b, 50, NewRcLifo(), benchA1R2)}
func BenchmarkRcLifo_A1R2x100(b *testing.B) { runParallelBench(b, 100, NewRcLifo(), benchA1R2)}
func BenchmarkRcLifo_A2R1x1(b *testing.B) { runParallelBench(b, 1, NewRcLifo(), benchA2R1)}
func BenchmarkRcLifo_A2R1x3(b *testing.B) { runParallelBench(b, 3, NewRcLifo(), benchA2R1)}
func BenchmarkRcLifo_A2R1x5(b *testing.B) { runParallelBench(b, 5, NewRcLifo(), benchA2R1)}
func BenchmarkRcLifo_A2R1x10(b *testing.B) { runParallelBench(b, 10, NewRcLifo(), benchA2R1)}
func BenchmarkRcLifo_A2R1x50(b *testing.B) { runParallelBench(b, 50, NewRcLifo(), benchA2R1)}
func BenchmarkRcLifo_A2R1x100(b *testing.B) { runParallelBench(b, 100, NewRcLifo(), benchA2R1)}

// RcFifo

func TestRcFifo(t *testing.T) {
	fifoTest(t, NewRcFifo())
	queuePTest(t, NewRcFifo())
	queueP2Test(t, NewRcFifo())
}
func BenchmarkRcFifo_ANRNx1(b *testing.B) { runParallelBench(b, 1, NewRcFifo(), benchANRN)}
func BenchmarkRcFifo_ANRNx3(b *testing.B) { runParallelBench(b, 3, NewRcFifo(), benchANRN)}
func BenchmarkRcFifo_ANRNx5(b *testing.B) { runParallelBench(b, 5, NewRcFifo(), benchANRN)}
func BenchmarkRcFifo_ANRNx10(b *testing.B) { runParallelBench(b, 10, NewRcFifo(), benchANRN)}
func BenchmarkRcFifo_ANRNx50(b *testing.B) { runParallelBench(b, 50, NewRcFifo(), benchANRN)}
func BenchmarkRcFifo_ANRNx100(b *testing.B) { runParallelBench(b, 100, NewRcFifo(), benchANRN)}
func BenchmarkRcFifo_A1R1x1(b *testing.B) { runParallelBench(b, 1, NewRcFifo(), benchA1R1)}
func BenchmarkRcFifo_A1R1x3(b *testing.B) { runParallelBench(b, 3, NewRcFifo(), benchA1R1)}
func BenchmarkRcFifo_A1R1x5(b *testing.B) { runParallelBench(b, 5, NewRcFifo(), benchA1R1)}
func BenchmarkRcFifo_A1R1x10(b *testing.B) { runParallelBench(b, 10, NewRcFifo(), benchA1R1)}
func BenchmarkRcFifo_A1R1x50(b *testing.B) { runParallelBench(b, 50, NewRcFifo(), benchA1R1)}
func BenchmarkRcFifo_A1R1x100(b *testing.B) { runParallelBench(b, 100, NewRcFifo(), benchA1R1)}
func BenchmarkRcFifo_A1R2x1(b *testing.B) { runParallelBench(b, 1, NewRcFifo(), benchA1R2)}
func BenchmarkRcFifo_A1R2x3(b *testing.B) { runParallelBench(b, 3, NewRcFifo(), benchA1R2)}
func BenchmarkRcFifo_A1R2x5(b *testing.B) { runParallelBench(b, 5, NewRcFifo(), benchA1R2)}
func BenchmarkRcFifo_A1R2x10(b *testing.B) { runParallelBench(b, 10, NewRcFifo(), benchA1R2)}
func BenchmarkRcFifo_A1R2x50(b *testing.B) { runParallelBench(b, 50, NewRcFifo(), benchA1R2)}
func BenchmarkRcFifo_A1R2x100(b *testing.B) { runParallelBench(b, 100, NewRcFifo(), benchA1R2)}
func BenchmarkRcFifo_A2R1x1(b *testing.B) { runParallelBench(b, 1, NewRcFifo(), benchA2R1)}
func BenchmarkRcFifo_A2R1x3(b *testing.B) { runParallelBench(b, 3, NewRcFifo(), benchA2R1)}
func BenchmarkRcFifo_A2R1x5(b *testing.B) { runParallelBench(b, 5, NewRcFifo(), benchA2R1)}
func BenchmarkRcFifo_A2R1x10(b *testing.B) { runParallelBench(b, 10, NewRcFifo(), benchA2R1)}
func BenchmarkRcFifo_A2R1x50(b *testing.B) { runParallelBench(b, 50, NewRcFifo(), benchA2R1)}
func BenchmarkRcFifo_A2R1x100(b *testing.B) { runParallelBench(b, 100, NewRcFifo(), benchA2R1)}

// RmLifo

func TestRmLifo(t *testing.T) {
	lifoTest(t, NewRmLifo())
	queuePTest(t, NewRmLifo())
	queueP2Test(t, NewRmLifo())
}
func BenchmarkRmLifo_ANRNx1(b *testing.B) { runParallelBench(b, 1, NewRmLifo(), benchANRN)}
func BenchmarkRmLifo_ANRNx3(b *testing.B) { runParallelBench(b, 3, NewRmLifo(), benchANRN)}
func BenchmarkRmLifo_ANRNx5(b *testing.B) { runParallelBench(b, 5, NewRmLifo(), benchANRN)}
func BenchmarkRmLifo_ANRNx10(b *testing.B) { runParallelBench(b, 10, NewRmLifo(), benchANRN)}
func BenchmarkRmLifo_ANRNx50(b *testing.B) { runParallelBench(b, 50, NewRmLifo(), benchANRN)}
func BenchmarkRmLifo_ANRNx100(b *testing.B) { runParallelBench(b, 100, NewRmLifo(), benchANRN)}
func BenchmarkRmLifo_A1R1x1(b *testing.B) { runParallelBench(b, 1, NewRmLifo(), benchA1R1)}
func BenchmarkRmLifo_A1R1x3(b *testing.B) { runParallelBench(b, 3, NewRmLifo(), benchA1R1)}
func BenchmarkRmLifo_A1R1x5(b *testing.B) { runParallelBench(b, 5, NewRmLifo(), benchA1R1)}
func BenchmarkRmLifo_A1R1x10(b *testing.B) { runParallelBench(b, 10, NewRmLifo(), benchA1R1)}
func BenchmarkRmLifo_A1R1x50(b *testing.B) { runParallelBench(b, 50, NewRmLifo(), benchA1R1)}
func BenchmarkRmLifo_A1R1x100(b *testing.B) { runParallelBench(b, 100, NewRmLifo(), benchA1R1)}
func BenchmarkRmLifo_A1R2x1(b *testing.B) { runParallelBench(b, 1, NewRmLifo(), benchA1R2)}
func BenchmarkRmLifo_A1R2x3(b *testing.B) { runParallelBench(b, 3, NewRmLifo(), benchA1R2)}
func BenchmarkRmLifo_A1R2x5(b *testing.B) { runParallelBench(b, 5, NewRmLifo(), benchA1R2)}
func BenchmarkRmLifo_A1R2x10(b *testing.B) { runParallelBench(b, 10, NewRmLifo(), benchA1R2)}
func BenchmarkRmLifo_A1R2x50(b *testing.B) { runParallelBench(b, 50, NewRmLifo(), benchA1R2)}
func BenchmarkRmLifo_A1R2x100(b *testing.B) { runParallelBench(b, 100, NewRmLifo(), benchA1R2)}
func BenchmarkRmLifo_A2R1x1(b *testing.B) { runParallelBench(b, 1, NewRmLifo(), benchA2R1)}
func BenchmarkRmLifo_A2R1x3(b *testing.B) { runParallelBench(b, 3, NewRmLifo(), benchA2R1)}
func BenchmarkRmLifo_A2R1x5(b *testing.B) { runParallelBench(b, 5, NewRmLifo(), benchA2R1)}
func BenchmarkRmLifo_A2R1x10(b *testing.B) { runParallelBench(b, 10, NewRmLifo(), benchA2R1)}
func BenchmarkRmLifo_A2R1x50(b *testing.B) { runParallelBench(b, 50, NewRmLifo(), benchA2R1)}
func BenchmarkRmLifo_A2R1x100(b *testing.B) { runParallelBench(b, 100, NewRmLifo(), benchA2R1)}

// RmFifo

func TestRmFifo(t *testing.T) {
	fifoTest(t, NewRmFifo())
	queuePTest(t, NewRmFifo())
	queueP2Test(t, NewRmFifo())
}
func BenchmarkRmFifo_ANRNx1(b *testing.B) { runParallelBench(b, 1, NewRmFifo(), benchANRN)}
func BenchmarkRmFifo_ANRNx3(b *testing.B) { runParallelBench(b, 3, NewRmFifo(), benchANRN)}
func BenchmarkRmFifo_ANRNx5(b *testing.B) { runParallelBench(b, 5, NewRmFifo(), benchANRN)}
func BenchmarkRmFifo_ANRNx10(b *testing.B) { runParallelBench(b, 10, NewRmFifo(), benchANRN)}
func BenchmarkRmFifo_ANRNx50(b *testing.B) { runParallelBench(b, 50, NewRmFifo(), benchANRN)}
func BenchmarkRmFifo_ANRNx100(b *testing.B) { runParallelBench(b, 100, NewRmFifo(), benchANRN)}
func BenchmarkRmFifo_A1R1x1(b *testing.B) { runParallelBench(b, 1, NewRmFifo(), benchA1R1)}
func BenchmarkRmFifo_A1R1x3(b *testing.B) { runParallelBench(b, 3, NewRmFifo(), benchA1R1)}
func BenchmarkRmFifo_A1R1x5(b *testing.B) { runParallelBench(b, 5, NewRmFifo(), benchA1R1)}
func BenchmarkRmFifo_A1R1x10(b *testing.B) { runParallelBench(b, 10, NewRmFifo(), benchA1R1)}
func BenchmarkRmFifo_A1R1x50(b *testing.B) { runParallelBench(b, 50, NewRmFifo(), benchA1R1)}
func BenchmarkRmFifo_A1R1x100(b *testing.B) { runParallelBench(b, 100, NewRmFifo(), benchA1R1)}
func BenchmarkRmFifo_A1R2x1(b *testing.B) { runParallelBench(b, 1, NewRmFifo(), benchA1R2)}
func BenchmarkRmFifo_A1R2x3(b *testing.B) { runParallelBench(b, 3, NewRmFifo(), benchA1R2)}
func BenchmarkRmFifo_A1R2x5(b *testing.B) { runParallelBench(b, 5, NewRmFifo(), benchA1R2)}
func BenchmarkRmFifo_A1R2x10(b *testing.B) { runParallelBench(b, 10, NewRmFifo(), benchA1R2)}
func BenchmarkRmFifo_A1R2x50(b *testing.B) { runParallelBench(b, 50, NewRmFifo(), benchA1R2)}
func BenchmarkRmFifo_A1R2x100(b *testing.B) { runParallelBench(b, 100, NewRmFifo(), benchA1R2)}
func BenchmarkRmFifo_A2R1x1(b *testing.B) { runParallelBench(b, 1, NewRmFifo(), benchA2R1)}
func BenchmarkRmFifo_A2R1x3(b *testing.B) { runParallelBench(b, 3, NewRmFifo(), benchA2R1)}
func BenchmarkRmFifo_A2R1x5(b *testing.B) { runParallelBench(b, 5, NewRmFifo(), benchA2R1)}
func BenchmarkRmFifo_A2R1x10(b *testing.B) { runParallelBench(b, 10, NewRmFifo(), benchA2R1)}
func BenchmarkRmFifo_A2R1x50(b *testing.B) { runParallelBench(b, 50, NewRmFifo(), benchA2R1)}
func BenchmarkRmFifo_A2R1x100(b *testing.B) { runParallelBench(b, 100, NewRmFifo(), benchA2R1)}
