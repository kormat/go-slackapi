package query

import (
	"bitbucket.org/kormaton/slapi/channel"
	"encoding/json"
	"github.com/golang/glog"
	"github.com/kr/pretty"
)

type Response struct {
	Ok      bool
	Warning *string
	Error   *string
	Channel *channel.Channel
}

func parse(data []byte) (Response, bool) {
	var r Response
	err := json.Unmarshal(data, &r)
	if err != nil {
		glog.Error("Unable to parse response: %v", err)
		return Response{}, false
	}
	if r.Error != nil {
		glog.Error("Error response from API: %s", *r.Error)
	}
	if r.Warning != nil {
		glog.Warning("Warning response from API: %s", *r.Warning)
	}
	return r, true
}

func (r *Response) String() string {
	return pretty.Sprint(r)
}
