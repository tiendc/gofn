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
