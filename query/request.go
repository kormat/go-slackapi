package query

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/util"
	"io/ioutil"
	"net/http"
	"net/url"
)

var client = &http.Client{}

func Request(cmd string, v url.Values) (Response, error) {
	u, err := createURL(cmd, v)
	if err != nil {
		return Response{}, err
	}
	glog.Infof("request: command: %#v", cmd)
	glog.V(2).Infof("request: url: %s", u)
	resp, err := client.Get(u.String())
	if err != nil {
		return Response{}, util.Error("query: http request error: %v", err)
	}
	glog.V(1).Infof("request: http response status: %v", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return Response{}, util.Error("query: error reading http response body: %v", err)
	}
	glog.V(1).Infof("request: http response body:\n%s", body)
	return Parse(body)
}

func createURL(cmd string, v url.Values) (*url.URL, error) {
	base := config.Cfg.APIEndpoint
	if len(base) == 0 {
		return nil, util.Error("API Endpoint url is empty")
	}
	v.Set("token", config.Cfg.Token)
	u, err := url.Parse(config.Cfg.APIEndpoint)
	if err != nil {
		return nil, util.Error("Unable to parse API endpoint url: %v", err)
	}
	u.Path = fmt.Sprintf("%s/%s", u.Path, cmd)
	u.RawQuery = v.Encode()
	return u, nil
}
