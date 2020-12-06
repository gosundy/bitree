package bitree

import (
	"math/rand"
	"testing"
)

const (
	BENCH_DATA_COUNT = 1000000
)

func BenchmarkBitNode_OneMillion(b *testing.B) {
	bitmap := NewBitTree()
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		j := uint32(0)
		for j = 0; j < BENCH_DATA_COUNT; j++ {
			err := bitmap.Set(j)
			if err != nil {
				b.Fatal(err)
			}
		}
	}
	b.StopTimer()

}
func BenchmarkBitNode_OneMillionWithRandom(b *testing.B) {
	bitmap := NewBitTree()
	datas := rand.Perm(BENCH_DATA_COUNT)
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, data := range datas {
			err := bitmap.Set(uint32(data))
			if err != nil {
				b.Fatal(err)
			}
		}
	}
	b.StopTimer()
}
func BenchmarkBitNode_OneMillionWithRandomOneHundredMillion(b *testing.B) {
	bitmap := NewBitTree()
	rand.Int31()
	datas := make([]int, BENCH_DATA_COUNT)
	for i := 0; i < BENCH_DATA_COUNT; i++ {
		datas[i] = rand.Intn(100000000)
	}
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, data := range datas {
			err := bitmap.Set(uint32(data))
			if err != nil {
				b.Fatal(err)
			}
		}
	}
	b.StopTimer()
}
