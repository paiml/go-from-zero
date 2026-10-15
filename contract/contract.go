// Package contract enforces a named runtime contract and prints its proof line.
package contract

import (
	"fmt"
	"io"
)

// Violation is the error returned when a contract does not hold.
type Violation struct{ Name string }

func (v Violation) Error() string { return "contract: " + v.Name + " VIOLATED" }

// Check prints "contract: <name> OK" to w when ok, else returns a Violation.
func Check(w io.Writer, name string, ok bool) error {
	if !ok {
		return Violation{Name: name}
	}
	_, err := fmt.Fprintf(w, "contract: %s OK\n", name)
	return err
}
