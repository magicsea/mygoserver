package console

import (
	"testing"
)

func Benchmark_newconsole(b *testing.B) {
	Console(nil)
	for i := 0; i < b.N; i++ {

	}
}
