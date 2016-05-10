package query

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/kr/pretty"
)

type Response struct {
	Ok       bool
	Warning  *string
	Error    *string
	Channel  *json.RawMessage
	Channels *json.RawMessage
}

/* bool return is for json decoding */
func Parse(data []byte) (Response, bool) {
	var r Response
	err := json.Unmarshal(data, &r)
	if err != nil {
		glog.Errorf("Unable to parse response: %v", err)
		return Response{}, false
	}
	if r.Error != nil {
		glog.Errorf("Error response from API: %s", *r.Error)
	}
	if r.Warning != nil {
		glog.Warningf("Warning response from API: %s", *r.Warning)
	}
	return r, true
}

func (r *Response) String() string {
	return pretty.Sprint(r)
}