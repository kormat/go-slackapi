package main

import (
	"github.com/golang/glog"
	"os"
)

func main() {
	ret := 0
	if _, err := parser.Parse(); err != nil {
		ret = 1
	}
	glog.Info("Exiting:", ret)
	glog.Flush()
	os.Exit(ret)
}
