// Package stats computes simple summaries of a slice of numbers.
package stats

import (
	"errors"
	"sort"
)

// ErrEmpty is returned when the input has no values.
var ErrEmpty = errors.New("stats: no values")

// Mean returns the arithmetic mean of xs.
func Mean(xs []float64) (float64, error) {
	if len(xs) == 0 {
		return 0, ErrEmpty
	}
	sum := 0.0
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs)), nil
}

// Median returns the middle value of xs; xs is not modified.
func Median(xs []float64) (float64, error) {
	if len(xs) == 0 {
		return 0, ErrEmpty
	}
	s := append([]float64(nil), xs...)
	sort.Float64s(s)
	mid := len(s) / 2
	if len(s)%2 == 1 {
		return s[mid], nil
	}
	return (s[mid-1] + s[mid]) / 2, nil
}
