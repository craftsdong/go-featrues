package bench

import (
	"strconv"
	"testing"
)

func BenchmarkMapRaw(b *testing.B) {
	b.StartTimer()
	m := make(map[string]int)
	for i := 0; i < b.N; i++ {
		m[strconv.Itoa(i)] = i
	}
	b.StopTimer()
}

func BenchmarkMapCap(b *testing.B) {
	b.StartTimer()
	m := make(map[string]int, b.N)
	for i := 0; i < b.N; i++ {
		m[strconv.Itoa(i)] = i
	}
	b.StopTimer()
}

func BenchmarkSliceRaw(b *testing.B) {
	b.StartTimer()
	var s []int
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
	b.StopTimer()
}

func BenchmarkSliceCapAppend(b *testing.B) {
	b.StartTimer()
	s := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
	b.StopTimer()
}

func BenchmarkSliceCapSet(b *testing.B) {
	b.StartTimer()
	s := make([]int, b.N, b.N)
	for i := 0; i < b.N; i++ {
		s[i] = i
	}
	b.StopTimer()
}
