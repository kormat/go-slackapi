package cmd

import (
	"os"
)

func Parse() {
	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}
}
