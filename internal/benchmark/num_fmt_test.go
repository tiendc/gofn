package benchmark

import (
	"strings"
	"testing"

	"github.com/tiendc/gofn"
)

// Benchmark_NumberUngroup
// Result:
//
//	Benchmark_NumberUngroup/gofn.NumberUngroup
//	Benchmark_NumberUngroup/gofn.NumberUngroup-10         	 2349999	       480.3 ns/op
//	Benchmark_NumberUngroup/strings.ReplaceAll
//	Benchmark_NumberUngroup/strings.ReplaceAll-10         	 1516845	       790.1 ns/op
func Benchmark_NumberUngroup(b *testing.B) {
	groupingSep := byte(',')
	prepareData := func() []string {
		return []string{"123,4567", "1234,567", "123,4567,1234,567", "567,123,456", "123,456,712,34",
			"123,4567", "1234,567", "123,4567,1234,567", "567,123,456", "123,456,712,34",
			"123,4567", "1234,567", "123,4567,1234,567", "567,123,456", "123,456,712,34"}
	}

	b.Run("gofn.NumberFmtUngroup", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			s := prepareData()
			for _, ss := range s {
				gofn.NumberFmtUngroup(ss, groupingSep)
			}
		}
	})

	b.Run("strings.ReplaceAll", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			s := prepareData()
			for _, ss := range s {
				strings.ReplaceAll(ss, string(groupingSep), "")
			}
		}
	})
}
