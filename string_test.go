package bench

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkStringJoin(b *testing.B) {
	b.StartTimer()
	s := ""
	for i := 0; i < b.N; i++ {
		s += "a"
	}
	b.StopTimer()
}

func BenchmarkBytesBuffer(b *testing.B) {
	b.StartTimer()
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.WriteString("a")
	}
	_ = buffer.String()
	b.StopTimer()
}

func BenchmarkStringsBuilder(b *testing.B) {
	b.StartTimer()
	var str strings.Builder
	for i := 0; i < b.N; i++ {
		str.WriteString("a")
	}
	_ = str.String()
	b.StopTimer()
}
