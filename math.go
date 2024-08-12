package gofn

// Abs calculates absolute value of an integer.
// Ref: http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html.
// Note: if inputs MinInt64, the result is negative.
func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}
