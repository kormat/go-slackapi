package config

import (
	"encoding/json"
	"github.com/golang/glog"
	"io/ioutil"
	"net/url"
	"os"
)

var Config struct {
	Token       string
	APIEndpoint string
}

var defAPIEndpoint = "https://slack.com/api"

func Load(path string) {
	glog.V(1).Infof("config: Loading from %#v", path)
	input, err := ioutil.ReadFile(path)
	if err != nil {
		glog.Errorf("config: Error reading file: %v", err)
		os.Exit(1)
	}
	err = json.Unmarshal(input, &Config)
	if err != nil {
		glog.Errorf("config: JSON parsing failure: %v", err)
		os.Exit(1)
	}
	if len(Config.Token) == 0 {
		glog.Errorf("config: No token specified.")
		os.Exit(1)
	}
	if len(Config.APIEndpoint) == 0 {
		Config.APIEndpoint = defAPIEndpoint
	}
}

func MakeURLValues(values map[string]string) url.Values {
	v := url.Values{}
	v.Set("token", Config.Token)
	for key, val := range values {
		v.Set(key, val)
	}
	return v
}
