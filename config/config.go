package config

import (
	"errors"
	"flag"
	"fmt"
	"github.com/kormat/go-slackapi/util"
	"io/ioutil"
	"net/url"
)

type Config struct {
	Token       string
	APIEndpoint string
	GlogFlags   []string
}

var Cfg *Config
var CfgErr error
var defAPIEndpoint = "https://slack.com/api"
var configured = false

func Load(path string) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		CfgErr = errors.New(fmt.Sprintf("config: Error reading file: %v", err))
		return
	}
	Cfg = &Config{}
	if util.ParseJSON(input, Cfg) != nil {
		CfgErr = errors.New(fmt.Sprintf("config: JSON parsing failure: %v", err))
		return
	}
	flag.CommandLine.Parse(Cfg.GlogFlags)
	if len(Cfg.Token) == 0 {
		CfgErr = errors.New(fmt.Sprintf("config: No token specified."))
		return
	}
	if len(Cfg.APIEndpoint) == 0 {
		Cfg.APIEndpoint = defAPIEndpoint
	}
	configured = true
}

func Set(cfg *Config) {
	if len(cfg.Token) == 0 {
		panic("slackapi.config: No token specified.")
	}
	if len(cfg.APIEndpoint) == 0 {
		cfg.APIEndpoint = defAPIEndpoint
	}
	Cfg = cfg
	configured = true
}

func MakeURLValues(values map[string]string) url.Values {
	if !configured {
		panic("slackapi.config: slackapi not configured")
	}
	v := url.Values{}
	v.Set("token", Cfg.Token)
	v.Set("pretty", "1")
	for key, val := range values {
		v.Set(key, val)
	}
	return v
}
