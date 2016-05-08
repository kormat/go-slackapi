package config

import (
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"io/ioutil"
	"os"
)

var config = flag.String("c", "slapi.json", "Config file")
var conf *Config

type Config struct {
	Token string
}

func Load() *Config {
	if conf != nil {
		return conf
	}
	glog.V(1).Infof("Loading config from %#v", *config)
	input, err := ioutil.ReadFile(*config)
	if err != nil {
		glog.Errorf("Unable to open config file `%v`: %v", *config, err)
		os.Exit(1)
	}
	c := &Config{}
	err = json.Unmarshal(input, c)
	if err != nil {
		glog.Errorf("Unable to parse config: %v", err)
		os.Exit(1)
	}
	if len(c.Token) == 0 {
		glog.Errorf("Config token empty")
		os.Exit(1)
	}
	return c
}
