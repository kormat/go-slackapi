package query

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/kormat/go-slackapi/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var client = &http.Client{}

func Request(cmd string, v url.Values) (Response, bool) {
	u := CreateURL(cmd, v)
	glog.Infof("request: command: %#v", cmd)
	glog.V(2).Infof("request: url: %s", u)
	if v.Get("token") == "" {
		glog.Errorf("query: No token specified")
		os.Exit(1)
	}
	resp, err := client.Get(u.String())
	glog.V(1).Infof("request: http response status: %v", resp.Status)
	if err != nil {
		return Response{}, false
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return Response{}, false
	}
	glog.V(1).Infof("request: http response body:\n%s", body)
	r, ok := Parse(body)
	return r, ok
}

func CreateURL(cmd string, v url.Values) *url.URL {
	base := config.Cfg.APIEndpoint
	if len(base) == 0 {
		glog.Error("API Endpoint url is empty")
		os.Exit(1)
	}
	u, err := url.Parse(config.Cfg.APIEndpoint)
	if err != nil {
		glog.Error("Unable to parse API endpoint url: %v", err)
		return nil
	}
	u.Path = fmt.Sprintf("%s/%s", u.Path, cmd)
	u.RawQuery = v.Encode()
	return u
}
