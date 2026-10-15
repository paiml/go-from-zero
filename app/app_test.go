package app

import (
	"bytes"
	"strings"
	"testing"
	"testing/iotest"
)

func run(in string, args ...string) (int, string, string) {
	var o, e bytes.Buffer
	code := Run(args, strings.NewReader(in), &o, &e)
	return code, o.String(), e.String()
}

func TestRun(t *testing.T) {
	cases := []struct {
		name string
		in   string
		args []string
		code int
		out  string
		err  string
	}{
		{"greet", "", []string{"greet", "Ada"}, 0, "Hello, Ada!\ncontract: greeting-is-nonempty-and-addressed OK\n", ""},
		{"greet arity", "", []string{"greet"}, 1, "", "gofz: greet takes exactly one NAME\n"},
		{"greet blank", "", []string{"greet", " "}, 1, "", "gofz: greet: name is empty\n"},
		{"stats", "", []string{"stats", "1", "2", "6"}, 0, "mean=3 median=2\ncontract: mean-and-median-within-min-max OK\n", ""},
		{"stats parse", "", []string{"stats", "x"}, 1, "", "gofz: strconv.ParseFloat: parsing \"x\": invalid syntax\n"},
		{"stats empty", "", []string{"stats"}, 1, "", "gofz: stats: no values\n"},
		{"wc", "a b\nc\n", []string{"wc"}, 0, "2 3 6\ncontract: words-never-exceed-bytes OK\n", ""},
		{"no args", "", nil, 2, "", usage + "\n"},
		{"unknown", "", []string{"nope"}, 2, "", usage + "\n"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			code, out, err := run(c.in, c.args...)
			if code != c.code || out != c.out || err != c.err {
				t.Fatalf("got %d %q %q; want %d %q %q", code, out, err, c.code, c.out, c.err)
			}
		})
	}
}

func TestRunWCReadError(t *testing.T) {
	var o, e bytes.Buffer
	if code := Run([]string{"wc"}, iotest.ErrReader(iotest.ErrTimeout), &o, &e); code != 1 {
		t.Fatalf("code = %d", code)
	}
}
