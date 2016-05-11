package config

import (
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"io/ioutil"
	"net/url"
	"os"
)

type Config struct {
	Token       string
	APIEndpoint string
	GlogFlags   []string
}

var Cfg *Config

var defAPIEndpoint = "https://slack.com/api"

func Load(path string) {
	glog.V(1).Infof("config: Loading from %#v", path)
	input, err := ioutil.ReadFile(path)
	if err != nil {
		glog.Errorf("config: Error reading file: %v", err)
		os.Exit(1)
	}
	Cfg = &Config{}
	err = json.Unmarshal(input, Cfg)
	if err != nil {
		glog.Errorf("config: JSON parsing failure: %v", err)
		os.Exit(1)
	}
	if len(Cfg.Token) == 0 {
		glog.Errorf("config: No token specified.")
		os.Exit(1)
	}
	if len(Cfg.APIEndpoint) == 0 {
		Cfg.APIEndpoint = defAPIEndpoint
	}
	flag.CommandLine.Parse(Cfg.GlogFlags)
}

func MakeURLValues(values map[string]string) url.Values {
	v := url.Values{}
	v.Set("token", Cfg.Token)
	for key, val := range values {
		v.Set(key, val)
	}
	return v
}
