// Package app is the gofz command line: greet, stats and wc.
package app

import (
	"fmt"
	"io"
	"strconv"

	"github.com/paiml/go-from-zero/contract"
	"github.com/paiml/go-from-zero/greet"
	"github.com/paiml/go-from-zero/stats"
	"github.com/paiml/go-from-zero/wc"
)

const usage = "usage: gofz greet NAME | stats N... | wc"

// Run executes one command and returns the process exit code.
func Run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, usage)
		return 2
	}
	var err error
	switch args[0] {
	case "greet":
		err = runGreet(args[1:], stdout)
	case "stats":
		err = runStats(args[1:], stdout)
	case "wc":
		err = runWC(stdin, stdout)
	default:
		fmt.Fprintln(stderr, usage)
		return 2
	}
	if err != nil {
		fmt.Fprintln(stderr, "gofz:", err)
		return 1
	}
	return 0
}

func runGreet(args []string, out io.Writer) error {
	if len(args) != 1 {
		return fmt.Errorf("greet takes exactly one NAME")
	}
	msg, err := greet.Greeting(args[0])
	if err != nil {
		return err
	}
	fmt.Fprintln(out, msg)
	return contract.Check(out, "greeting-is-nonempty-and-addressed", len(msg) > len("Hello, !"))
}

func runStats(args []string, out io.Writer) error {
	xs := make([]float64, 0, len(args))
	for _, a := range args {
		x, err := strconv.ParseFloat(a, 64)
		if err != nil {
			return err
		}
		xs = append(xs, x)
	}
	mean, err := stats.Mean(xs)
	if err != nil {
		return err
	}
	med, _ := stats.Median(xs)
	fmt.Fprintf(out, "mean=%g median=%g\n", mean, med)
	lo, hi := xs[0], xs[0]
	for _, x := range xs {
		lo, hi = min(lo, x), max(hi, x)
	}
	return contract.Check(out, "mean-and-median-within-min-max", lo <= mean && mean <= hi && lo <= med && med <= hi)
}

func runWC(in io.Reader, out io.Writer) error {
	c, err := wc.Count(in)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%d %d %d\n", c.Lines, c.Words, c.Bytes)
	return contract.Check(out, "words-never-exceed-bytes", c.Words <= c.Bytes)
}
