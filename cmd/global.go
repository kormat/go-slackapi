package cmd

import (
	"bitbucket.org/kormaton/slapi/config"
	"flag"
	"github.com/jessevdk/go-flags"
)

type GlobalFlags struct {
	Config func(string) `short:"c" long:"config" description:"Config file" default:"slapi.json"`
}

var globalFlags GlobalFlags
var parser = flags.NewParser(&globalFlags, flags.Default)

func init() {
	globalFlags.Config = func(path string) {
		config.Load(path)
	}
	// FIXME(kormat): glog expects flag.Parse to be called, but that overrides
	// go-flags, so for now just bypass flag.Parse
	flag.CommandLine.Parse([]string{"-alsologtostderr"})
}
