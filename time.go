package gofn

import "time"

// MinTime finds the minimum time in the list.
// NOTE: if zero time is in the list, the result will be zero.
func MinTime(t1 time.Time, s ...time.Time) time.Time {
	min := t1
	for i := range s {
		if s[i].Before(min) {
			min = s[i]
		}
	}
	return min
}

// MaxTime finds the maximum time in the list
func MaxTime(t1 time.Time, s ...time.Time) time.Time {
	max := t1
	for i := range s {
		if s[i].After(max) {
			max = s[i]
		}
	}
	return max
}

// MinMaxTime gets the minimum and maximum time values in the list
func MinMaxTime(t1 time.Time, s ...time.Time) (time.Time, time.Time) {
	minTime := t1
	maxTime := t1
	for i := range s {
		if s[i].Before(minTime) {
			minTime = s[i]
		} else if s[i].After(maxTime) {
			maxTime = s[i]
		}
	}
	return minTime, maxTime
}

// ExecDuration measures time duration of running a function
func ExecDuration(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

// ExecDuration1 measures time duration of running a function
func ExecDuration1[T any](fn func() T) (T, time.Duration) {
	start := time.Now()
	val := fn()
	return val, time.Since(start)
}

// ExecDuration2 measures time duration of running a function
func ExecDuration2[T1, T2 any](fn func() (T1, T2)) (T1, T2, time.Duration) {
	start := time.Now()
	val1, val2 := fn()
	return val1, val2, time.Since(start)
}

// ExecDuration3 measures time duration of running a function
func ExecDuration3[T1, T2, T3 any](fn func() (T1, T2, T3)) (T1, T2, T3, time.Duration) {
	start := time.Now()
	val1, val2, val3 := fn()
	return val1, val2, val3, time.Since(start)
}

// ExecDelay is an alias of time.AfterFunc
var ExecDelay = time.AfterFunc
