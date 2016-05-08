package cmd

import (
	"fmt"
	"github.com/golang/glog"
	"strings"
)

func paramCheck(cmd string, want, got []string) bool {
	if len(want) != len(got) {
		desc := ""
		if len(want) > 0 {
			desc = fmt.Sprintf(" (%s)", strings.Join(want, ", "))
		}
		glog.Errorf("%s: got %d arg(s), requires %d arg(s)%s", cmd, len(got), len(want), desc)
		return false
	}
	return true
}
