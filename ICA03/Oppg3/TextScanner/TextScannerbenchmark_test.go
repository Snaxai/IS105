package textscanner

import (
	"math/rand"
	"testing"
	"time"
)

var pg100 = "pg100.txt"
var fil10mb = "eks2.txt"
var fil20mb = "eks1.txt"

//var storfil = "storfil.txt"

func init() {
	seed := time.Now().Unix()
	rand.Seed(seed)
}

func perm(n int) (out []int) {
	for _, v := range rand.Perm(n) {
		out = append(out, v)
	}
	return
}

func BenchmarkFinnBufferLiten(b *testing.B) {
	FinnBuffer(pg100, b)
}
func BenchmarkFinnBufferMiddels(b *testing.B) {
	FinnBuffer(fil10mb, b)
}

func BenchmarkFinnBufferStor(b *testing.B) {
	FinnBuffer(fil20mb, b)
}

func BenchmarkFinnBufferStorfil(b *testing.B) {
	//FinnBuffer(storfil, b)
	//storfil skulle egentlig vært en stor fil på 200mb men kunne ikke legge det opp til repo
}

func FinnBuffer(i string, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		b.StartTimer()
		FinnAlleTest(i)

	}
}
