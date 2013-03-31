package main
// +build ignore

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

const Header = `package goqueuestest

import "testing"

`

const tests = `
func Test{{.Name}}(t *testing.T) {
	{{.Type}}Test(t, {{.Ctor}}())
	queuePTest(t, {{.Ctor}}())
	queueP2Test(t, {{.Ctor}}())
}

func Benchmark{{.Name}}(b *testing.B) {
	queueBench(b, {{.Ctor}}())
}

func Benchmark{{.Name}}P(b *testing.B) {
	queueBenchP(b, {{.Ctor}}())
}

func Benchmark{{.Name}}P2(b *testing.B) {
	queueBenchP2(b, {{.Ctor}}())
}
`

const testsSized = `
func Test{{.Name}}(t *testing.T) {
	{{.Type}}Test(t, {{.Ctor}}(testItemCount))
	queuePTest(t, {{.Ctor}}(testItemCount))
	queueP2Test(t, {{.Ctor}}(testItemCount))
}

func Benchmark{{.Name}}S(b *testing.B) {
	queueBench(b, {{.Ctor}}(b.N))
}

func Benchmark{{.Name}}P(b *testing.B) {
	queueBenchP(b, {{.Ctor}}(b.N))
}

func Benchmark{{.Name}}P2(b *testing.B) {
	queueBenchP2(b, {{.Ctor}}(b.N))
}
`

var queues = map[string]string {
	"CFifo" : "NewChanFifo",
	"LcFifo" : "NewListCFifo",
	"LcLifo" : "NewListCLifo",
	"LmFifo" : "NewListFifo",
	"LmLifo" : "NewListLifo",
	"ZFifo" : "NewZFifo",
	"ZLifo" : "NewZLifo",
	"ZcFifo" : "NewZFifoFreechan",
	"ZrFifo" : "NewZFifoFreering",
	"RmLifo" : "NewRLifo",
	"RmFifo" : "NewRFifo",
	"SmLifo" : "NewSLifo",
	"SmFifo" : "NewSFifo",
}

type queue struct { Name, Ctor, Type string }

var sizedQueues = map[string] bool {
	"CFifo" : true,
}

func main() {
	simple := template.Must(template.New("simple").Parse(tests))
	sized := template.Must(template.New("sized").Parse(testsSized))

	fmt.Printf(Header)
	for name, ctor := range(queues) {
		t := simple
		if sizedQueues[name] {
			t = sized
		}

		typ := "fifo"
		if strings.Contains(name, "Lifo") {
			typ = "lifo"
		}

		t.Execute(os.Stdout, &queue{name, ctor, typ})
	}
}