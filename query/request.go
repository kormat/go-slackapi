package query

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"net/url"
)

var base_url = flag.String("query.url", "https://slack.com/api", "Remote API endpoint")

var tr = &http.Transport{}
var client = &http.Client{Transport: tr}

type params map[string]string

func Request(cmd string, v url.Values) (Response, bool) {
	u := CreateURL(cmd, v)
	glog.Infof("Request: %#v", u.String())
	resp, err := client.Get(u.String())
	if err != nil {
		return Response{}, false
	}
	json, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return Response{}, false
	}
	r, ok := Parse(json)
	return r, ok
}

func CreateURL(cmd string, v url.Values) *url.URL {
	u, err := url.Parse(*base_url)
	if err != nil {
		glog.Error("Unable to parse API endpoint url: %v", err)
		return nil
	}
	u.Path = fmt.Sprintf("%s/%s", u.Path, cmd)
	u.RawQuery = v.Encode()
	return u
}
