// Command gofz is the course-one demo CLI.
package main

import (
	"os"

	"github.com/paiml/go-from-zero/app"
)

var exit = os.Exit

func main() {
	exit(app.Run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
