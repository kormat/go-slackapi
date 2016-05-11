package config

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
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

func Load(path string) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		CfgErr = errors.New(fmt.Sprintf("config: Error reading file: %v", err))
		return
	}
	Cfg = &Config{}
	err = json.Unmarshal(input, Cfg)
	if err != nil {
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
}

func MakeURLValues(values map[string]string) url.Values {
	v := url.Values{}
	v.Set("token", Cfg.Token)
	v.Set("pretty", "1")
	for key, val := range values {
		v.Set(key, val)
	}
	return v
}
