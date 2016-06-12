package query

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/kormat/go-slackapi/util"
	"github.com/kr/pretty"
)

type Response struct {
	// Meta data
	Ok      bool
	Warning *string
	Error   *string
	CacheTS *int `json:"cache_ts"`
	// Body
	Channel  *json.RawMessage
	Channels *json.RawMessage
	Group    *json.RawMessage
	Groups   *json.RawMessage
	Members  *json.RawMessage
	User     *json.RawMessage
}

func parseResponse(data []byte) (Response, error) {
	var r Response
	err := json.Unmarshal(data, &r)
	if err != nil {
		return Response{}, util.ErrorLog("query: unable to parse response json: %v\n%s", err, data)
	}
	if r.Error != nil {
		err = util.ErrorLog("query: api error: %s", *r.Error)
	}
	if r.Warning != nil {
		glog.Warningf("query: api warning: %s", *r.Warning)
	}
	return r, err
}

func (r *Response) String() string {
	return pretty.Sprint(r)
}
