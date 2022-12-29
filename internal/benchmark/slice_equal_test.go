package benchmark

import (
	"reflect"
	"testing"

	"github.com/tiendc/gofn"
	"github.com/tiendc/gofn/internal/testdata"
)

// Benchmark_Slice_Equal
// Result:
//
//	Benchmark_Slice_Equal/StructSlice/Equal
//	Benchmark_Slice_Equal/StructSlice/Equal-8         	510845715	         2.047 ns/op
//	Benchmark_Slice_Equal/StructSlice/ContentEqual
//	Benchmark_Slice_Equal/StructSlice/ContentEqual-8  	583167950	         2.061 ns/op
//	Benchmark_Slice_Equal/StructSlice/DeepEqual
//	Benchmark_Slice_Equal/StructSlice/DeepEqual-8     	15403771	         79.19 ns/op
//	Benchmark_Slice_Equal/IntSlice/Equal
//	Benchmark_Slice_Equal/IntSlice/Equal-8            	589706185	         2.087 ns/op
//	Benchmark_Slice_Equal/IntSlice/ContentEqual
//	Benchmark_Slice_Equal/IntSlice/ContentEqual-8     	523120755	         2.194 ns/op
//	Benchmark_Slice_Equal/IntSlice/DeepEqual
//	Benchmark_Slice_Equal/IntSlice/DeepEqual-8        	15243183	         77.93 ns/op
func Benchmark_Slice_Equal(b *testing.B) {
	srcStructSlice := testdata.BigStructSlice
	dstStructSlice := make([]testdata.BigStruct, 0, len(srcStructSlice))
	copy(dstStructSlice, srcStructSlice)

	b.Run("StructSlice/Equal", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gofn.Equal(srcStructSlice, dstStructSlice)
		}
	})

	b.Run("StructSlice/ContentEqual", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gofn.ContentEqual(srcStructSlice, dstStructSlice)
		}
	})

	b.Run("StructSlice/DeepEqual", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			reflect.DeepEqual(srcStructSlice, dstStructSlice)
		}
	})

	srcInt := testdata.IntSlice
	dstInt := make([]int, 0, len(srcInt))
	copy(dstInt, srcInt)

	b.Run("IntSlice/Equal", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gofn.Equal(srcInt, dstInt)
		}
	})

	b.Run("IntSlice/ContentEqual", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gofn.ContentEqual(srcInt, dstInt)
		}
	})

	b.Run("IntSlice/DeepEqual", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			reflect.DeepEqual(srcInt, dstInt)
		}
	})
}
