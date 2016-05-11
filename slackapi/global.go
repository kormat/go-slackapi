package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/kormat/go-slackapi/config"
)

type GlobalFlags struct {
	Config func(string) `short:"c" long:"config" description:"Config file" default:"slackapi.json"`
}

var globalFlags GlobalFlags
var parser = flags.NewParser(&globalFlags, flags.Default)

func init() {
	globalFlags.Config = func(path string) {
		config.Load(path)
	}
}
