// Package wc counts lines, words and bytes in a stream.
package wc

import (
	"bufio"
	"io"
	"strings"
)

// Counts holds the three totals.
type Counts struct {
	Lines, Words, Bytes int
}

// Count reads r to the end and returns its totals.
func Count(r io.Reader) (Counts, error) {
	var c Counts
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadString('\n')
		c.Bytes += len(line)
		c.Words += len(strings.Fields(line))
		if strings.HasSuffix(line, "\n") {
			c.Lines++
		}
		if err == io.EOF {
			return c, nil
		}
		if err != nil {
			return c, err
		}
	}
}
