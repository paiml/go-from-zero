package wc

import (
	"strings"
	"testing"
	"testing/iotest"
)

func TestCount(t *testing.T) {
	got, err := Count(strings.NewReader("one two\nthree\nfour"))
	want := Counts{Lines: 2, Words: 4, Bytes: 18}
	if err != nil || got != want {
		t.Fatalf("Count = %+v, %v; want %+v", got, err, want)
	}
}

func TestCountReadError(t *testing.T) {
	if _, err := Count(iotest.ErrReader(iotest.ErrTimeout)); err != iotest.ErrTimeout {
		t.Fatalf("err = %v; want ErrTimeout", err)
	}
}
