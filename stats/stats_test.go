package stats

import "testing"

func TestMean(t *testing.T) {
	if _, err := Mean(nil); err != ErrEmpty {
		t.Fatalf("empty: err = %v", err)
	}
	got, err := Mean([]float64{1, 2, 3, 6})
	if err != nil || got != 3 {
		t.Fatalf("Mean = %v, %v; want 3", got, err)
	}
}

func TestMedian(t *testing.T) {
	if _, err := Median(nil); err != ErrEmpty {
		t.Fatalf("empty: err = %v", err)
	}
	in := []float64{5, 1, 3}
	got, _ := Median(in)
	if got != 3 {
		t.Fatalf("odd Median = %v; want 3", got)
	}
	if in[0] != 5 {
		t.Fatalf("Median modified its input: %v", in)
	}
	got, _ = Median([]float64{4, 1, 3, 2})
	if got != 2.5 {
		t.Fatalf("even Median = %v; want 2.5", got)
	}
}
