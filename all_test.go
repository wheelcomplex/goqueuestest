package goqueuestest
// DO NOT EDIT, GENERATED CODE
import "testing"


func TestCFifo(t *testing.T) {
	fifoTest(t, NewChanFifo(testItemCount))
	queuePTest(t, NewChanFifo(testItemCount))
	queueP2Test(t, NewChanFifo(testItemCount))
}

func BenchmarkCFifoSing(b *testing.B) {
	queueBenchSing(b, NewChanFifo(b.N))
}

func BenchmarkCFifoANRN(b *testing.B) {
	queueBenchANRN(b, NewChanFifo(b.N))
}

func BenchmarkCFifoA1R1(b *testing.B) {
	queueBenchA1R1(b, NewChanFifo(b.N))
}

func BenchmarkCFifoA2R1(b *testing.B) {
	queueBenchA2R1(b, NewChanFifo(b.N))
}

func BenchmarkCFifoA1R2(b *testing.B) {
	queueBenchA1R2(b, NewChanFifo(b.N))
}

func TestLcLifo(t *testing.T) {
	lifoTest(t, NewListCLifo())
	queuePTest(t, NewListCLifo())
	queueP2Test(t, NewListCLifo())
}

func BenchmarkLcLifoSing(b *testing.B) {
	queueBenchSing(b, NewListCLifo())
}

func BenchmarkLcLifoANRN(b *testing.B) {
	queueBenchANRN(b, NewListCLifo())
}

func BenchmarkLcLifoA1R1(b *testing.B) {
	queueBenchA1R1(b, NewListCLifo())
}

func BenchmarkLcLifoA2R1(b *testing.B) {
	queueBenchA2R1(b, NewListCLifo())
}
func BenchmarkLcLifoA1R2(b *testing.B) {
	queueBenchA1R2(b, NewListCLifo())
}

func TestLcFifo(t *testing.T) {
	fifoTest(t, NewListCFifo())
	queuePTest(t, NewListCFifo())
	queueP2Test(t, NewListCFifo())
}

func BenchmarkLcFifoSing(b *testing.B) {
	queueBenchSing(b, NewListCFifo())
}

func BenchmarkLcFifoANRN(b *testing.B) {
	queueBenchANRN(b, NewListCFifo())
}

func BenchmarkLcFifoA1R1(b *testing.B) {
	queueBenchA1R1(b, NewListCFifo())
}

func BenchmarkLcFifoA2R1(b *testing.B) {
	queueBenchA2R1(b, NewListCFifo())
}
func BenchmarkLcFifoA1R2(b *testing.B) {
	queueBenchA1R2(b, NewListCFifo())
}

func TestLmLifo(t *testing.T) {
	lifoTest(t, NewListLifo())
	queuePTest(t, NewListLifo())
	queueP2Test(t, NewListLifo())
}

func BenchmarkLmLifoSing(b *testing.B) {
	queueBenchSing(b, NewListLifo())
}

func BenchmarkLmLifoANRN(b *testing.B) {
	queueBenchANRN(b, NewListLifo())
}

func BenchmarkLmLifoA1R1(b *testing.B) {
	queueBenchA1R1(b, NewListLifo())
}

func BenchmarkLmLifoA2R1(b *testing.B) {
	queueBenchA2R1(b, NewListLifo())
}
func BenchmarkLmLifoA1R2(b *testing.B) {
	queueBenchA1R2(b, NewListLifo())
}

func TestLmFifo(t *testing.T) {
	fifoTest(t, NewListFifo())
	queuePTest(t, NewListFifo())
	queueP2Test(t, NewListFifo())
}

func BenchmarkLmFifoSing(b *testing.B) {
	queueBenchSing(b, NewListFifo())
}

func BenchmarkLmFifoANRN(b *testing.B) {
	queueBenchANRN(b, NewListFifo())
}

func BenchmarkLmFifoA1R1(b *testing.B) {
	queueBenchA1R1(b, NewListFifo())
}

func BenchmarkLmFifoA2R1(b *testing.B) {
	queueBenchA2R1(b, NewListFifo())
}
func BenchmarkLmFifoA1R2(b *testing.B) {
	queueBenchA1R2(b, NewListFifo())
}

func TestZLifo(t *testing.T) {
	lifoTest(t, NewZLifo())
	queuePTest(t, NewZLifo())
	queueP2Test(t, NewZLifo())
}

func BenchmarkZLifoSing(b *testing.B) {
	queueBenchSing(b, NewZLifo())
}

func BenchmarkZLifoANRN(b *testing.B) {
	queueBenchANRN(b, NewZLifo())
}

func BenchmarkZLifoA1R1(b *testing.B) {
	queueBenchA1R1(b, NewZLifo())
}

func BenchmarkZLifoA2R1(b *testing.B) {
	queueBenchA2R1(b, NewZLifo())
}
func BenchmarkZLifoA1R2(b *testing.B) {
	queueBenchA1R2(b, NewZLifo())
}

func TestZFifo(t *testing.T) {
	fifoTest(t, NewZFifo())
	queuePTest(t, NewZFifo())
	queueP2Test(t, NewZFifo())
}

func BenchmarkZFifoSing(b *testing.B) {
	queueBenchSing(b, NewZFifo())
}

func BenchmarkZFifoANRN(b *testing.B) {
	queueBenchANRN(b, NewZFifo())
}

func BenchmarkZFifoA1R1(b *testing.B) {
	queueBenchA1R1(b, NewZFifo())
}

func BenchmarkZFifoA2R1(b *testing.B) {
	queueBenchA2R1(b, NewZFifo())
}
func BenchmarkZFifoA1R2(b *testing.B) {
	queueBenchA1R2(b, NewZFifo())
}

func TestZcFifo(t *testing.T) {
	fifoTest(t, NewZFifoFreechan())
	queuePTest(t, NewZFifoFreechan())
	queueP2Test(t, NewZFifoFreechan())
}

func BenchmarkZcFifoSing(b *testing.B) {
	queueBenchSing(b, NewZFifoFreechan())
}

func BenchmarkZcFifoANRN(b *testing.B) {
	queueBenchANRN(b, NewZFifoFreechan())
}

func BenchmarkZcFifoA1R1(b *testing.B) {
	queueBenchA1R1(b, NewZFifoFreechan())
}

func BenchmarkZcFifoA2R1(b *testing.B) {
	queueBenchA2R1(b, NewZFifoFreechan())
}
func BenchmarkZcFifoA1R2(b *testing.B) {
	queueBenchA1R2(b, NewZFifoFreechan())
}

func TestZrFifo(t *testing.T) {
	fifoTest(t, NewZFifoFreering())
	queuePTest(t, NewZFifoFreering())
	queueP2Test(t, NewZFifoFreering())
}

func BenchmarkZrFifoSing(b *testing.B) {
	queueBenchSing(b, NewZFifoFreering())
}

func BenchmarkZrFifoANRN(b *testing.B) {
	queueBenchANRN(b, NewZFifoFreering())
}

func BenchmarkZrFifoA1R1(b *testing.B) {
	queueBenchA1R1(b, NewZFifoFreering())
}

func BenchmarkZrFifoA2R1(b *testing.B) {
	queueBenchA2R1(b, NewZFifoFreering())
}
func BenchmarkZrFifoA1R2(b *testing.B) {
	queueBenchA1R2(b, NewZFifoFreering())
}

func TestRmLifo(t *testing.T) {
	lifoTest(t, NewRLifo())
	queuePTest(t, NewRLifo())
	queueP2Test(t, NewRLifo())
}

func BenchmarkRmLifoSing(b *testing.B) {
	queueBenchSing(b, NewRLifo())
}

func BenchmarkRmLifoANRN(b *testing.B) {
	queueBenchANRN(b, NewRLifo())
}

func BenchmarkRmLifoA1R1(b *testing.B) {
	queueBenchA1R1(b, NewRLifo())
}

func BenchmarkRmLifoA2R1(b *testing.B) {
	queueBenchA2R1(b, NewRLifo())
}
func BenchmarkRmLifoA1R2(b *testing.B) {
	queueBenchA1R2(b, NewRLifo())
}

func TestRmFifo(t *testing.T) {
	fifoTest(t, NewRFifo())
	queuePTest(t, NewRFifo())
	queueP2Test(t, NewRFifo())
}

func BenchmarkRmFifoSing(b *testing.B) {
	queueBenchSing(b, NewRFifo())
}

func BenchmarkRmFifoANRN(b *testing.B) {
	queueBenchANRN(b, NewRFifo())
}

func BenchmarkRmFifoA1R1(b *testing.B) {
	queueBenchA1R1(b, NewRFifo())
}

func BenchmarkRmFifoA2R1(b *testing.B) {
	queueBenchA2R1(b, NewRFifo())
}
func BenchmarkRmFifoA1R2(b *testing.B) {
	queueBenchA1R2(b, NewRFifo())
}

func TestSmLifo(t *testing.T) {
	lifoTest(t, NewSLifo())
	queuePTest(t, NewSLifo())
	queueP2Test(t, NewSLifo())
}

func BenchmarkSmLifoSing(b *testing.B) {
	queueBenchSing(b, NewSLifo())
}

func BenchmarkSmLifoANRN(b *testing.B) {
	queueBenchANRN(b, NewSLifo())
}

func BenchmarkSmLifoA1R1(b *testing.B) {
	queueBenchA1R1(b, NewSLifo())
}

func BenchmarkSmLifoA2R1(b *testing.B) {
	queueBenchA2R1(b, NewSLifo())
}
func BenchmarkSmLifoA1R2(b *testing.B) {
	queueBenchA1R2(b, NewSLifo())
}

func TestSmFifo(t *testing.T) {
	fifoTest(t, NewSFifo())
	queuePTest(t, NewSFifo())
	queueP2Test(t, NewSFifo())
}

func BenchmarkSmFifoSing(b *testing.B) {
	queueBenchSing(b, NewSFifo())
}

func BenchmarkSmFifoANRN(b *testing.B) {
	queueBenchANRN(b, NewSFifo())
}

func BenchmarkSmFifoA1R1(b *testing.B) {
	queueBenchA1R1(b, NewSFifo())
}

func BenchmarkSmFifoA2R1(b *testing.B) {
	queueBenchA2R1(b, NewSFifo())
}
func BenchmarkSmFifoA1R2(b *testing.B) {
	queueBenchA1R2(b, NewSFifo())
}
