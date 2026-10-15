package contract

import (
	"bytes"
	"errors"
	"testing"
)

type failWriter struct{}

func (failWriter) Write([]byte) (int, error) { return 0, errors.New("boom") }

func TestCheck(t *testing.T) {
	var b bytes.Buffer
	if err := Check(&b, "x", true); err != nil || b.String() != "contract: x OK\n" {
		t.Fatalf("got %q, %v", b.String(), err)
	}
	err := Check(&b, "y", false)
	if err == nil || err.Error() != "contract: y VIOLATED" {
		t.Fatalf("err = %v", err)
	}
	if Check(failWriter{}, "z", true) == nil {
		t.Fatal("write error swallowed")
	}
}
