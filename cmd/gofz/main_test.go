package main

import (
	"os"
	"testing"
)

func TestMain_Exit(t *testing.T) {
	code := -1
	exit = func(c int) { code = c }
	old := os.Args
	defer func() { os.Args = old; exit = os.Exit }()
	os.Args = []string{"gofz"}
	main()
	if code != 2 {
		t.Fatalf("exit code = %d; want 2", code)
	}
}
