package cmd

import (
	"bitbucket.org/kormaton/slapi/channels"
	"flag"
	"fmt"
	"os"
	"strings"
)

func Cmd() bool {
	args := flag.Args()
	cmd := args[0]
	args = args[1:]
	cmd_parts := strings.SplitN(cmd, ".", 2)
	switch cmd_parts[0] {
	case "channels":
		return cmdChannel(cmd, cmd_parts[1], args)
	}
	fmt.Fprintf(os.Stderr, "Unsupported command %#v\n", flag.Arg(0))
	return false
}

func cmdChannel(cmd, method string, params []string) bool {
	switch method {
	case "info":
		if !paramCheck(cmd, []string{"channel ID"}, params) {
			return false
		}
		c, ok := channels.Info(flag.Arg(1))
		if ok {
			fmt.Printf("%v\n", c)
		}
		return ok
	case "list":
		if !paramCheck(cmd, []string{}, params) {
			return false
		}
		chans, ok := channels.List()
		if !ok {
			fmt.Fprint(os.Stderr, "Failed to get channel list")
			return false
		}
		for i, c := range chans {
			fmt.Printf("%d. %s (Id: %s)\n", i, c.Name, c.Id)
		}
		return true
	}
	fmt.Fprintf(os.Stderr, "Unsupported channels method %#v\n", method)
	return false
}
