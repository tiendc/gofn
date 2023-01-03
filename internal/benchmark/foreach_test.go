package benchmark

import (
	"testing"

	"github.com/tiendc/gofn"
	"github.com/tiendc/gofn/internal/testdata"
)

// Benchmark_ForEach ForEach vs ForEachPtr on big struct slices
// Result:
//
//	Benchmark_ForEach/builtin_For-Index
//	Benchmark_ForEach/builtin_For-Index-8         	34310325	        33.90 ns/op
//	Benchmark_ForEach/builtin_ForEach
//	Benchmark_ForEach/builtin_ForEach-8           	36727526	        33.68 ns/op
//	Benchmark_ForEach/gofn.ForEach
//	Benchmark_ForEach/gofn.ForEach-8              	34559764	        34.38 ns/op
//	Benchmark_ForEach/gofn.ForEachPtr
//	Benchmark_ForEach/gofn.ForEachPtr-8           	97860658	        12.11 ns/op
func Benchmark_ForEach(b *testing.B) {
	slice := testdata.BigStructSlice
	fn := func(i int, t testdata.BigStruct) {
	}

	b.Run("builtin For-Index", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for i := 0; i < len(slice); i++ {
				fn(i, slice[i])
			}
		}
	})

	b.Run("builtin ForEach", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for i, v := range slice {
				fn(i, v)
			}
		}
	})

	b.Run("gofn.ForEach", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gofn.ForEach(slice, fn)
		}
	})

	b.Run("gofn.ForEachPtr", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gofn.ForEachPtr(slice, func(i int, t *testdata.BigStruct) {})
		}
	})
}
