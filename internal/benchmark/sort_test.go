package benchmark

import (
	"github.com/moneyforwardvietnam/gofn"
	"sort"
	"testing"
)

// Benchmark_Sort
// Result:
//
//	Benchmark_Sort/sort.Ints
//	Benchmark_Sort/sort.Ints-8         	 1186268	       920.1 ns/op
//	Benchmark_Sort/sort.Slice
//	Benchmark_Sort/sort.Slice-8        	 1268170	       948.1 ns/op
//	Benchmark_Sort/gofn.Sort
//	Benchmark_Sort/gofn.Sort-8         	 1256199	       967.0 ns/op
func Benchmark_Sort(b *testing.B) {
	prepareData := func() []int {
		return []int{-112312, 2312, -12314, 0, 0, 12, 1, 31238, -312545, -6456 - 23423,
			12312, 434, 545, 123, 7567, 3123, 23534, 45654, 0, 0, 1, 2, 3, 2, 1, 0, -123, -123, -123,
			-112312, 2312, -12314, 0, 0, 12, 1, 31238, -312545, -6456 - 23423,
			12312, 434, 545, 123, 7567, 3123, 23534, 45654, 0, 0, 1, 2, 3, 2, 1, 0, -123, -123, -123}
	}

	b.Run("sort.Ints", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			sort.Ints(prepareData())
		}
	})

	b.Run("sort.Slice", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			s := prepareData()
			sort.Slice(s, func(i, j int) bool {
				return s[i] < s[j]
			})
		}
	})

	b.Run("gofn.Sort", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gofn.Sort(prepareData())
		}
	})
}
