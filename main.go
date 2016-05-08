package main

import (
	"bitbucket.org/kormaton/slapi/cmd"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "slapi: No command given\n")
		os.Exit(1)
	}
	if !cmd.Cmd() {
		os.Exit(1)
	}
}
