package algorithms

import (
	"math/rand"
	"testing"
	"time"
)

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

func BenchmarkBSortM100(b *testing.B) {
	benchmarkBSortM(100, b)
}

func BenchmarkBSortM1000(b *testing.B) {
	benchmarkBSortM(1000, b)
}

func BenchmarkBSortM10000(b *testing.B) {
	benchmarkBSortM(10000, b)
}

func benchmarkBSortM(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		Bubblesort(values)
	}
}
