package groups

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/kormat/go-slackapi/channels"
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/query"
	"github.com/kr/pretty"
)

type GroupInfo struct {
	Id         string
	Name       string
	IsGroup    bool `json:"is_group"`
	Created    int
	Creator    string
	IsArchived bool `json:"is_archived"`
	IsMpim     bool `json:"is_mpim"`
	Members    []string
	Topic      channels.TopicPurpose
	Purpose    channels.TopicPurpose
}

func Info(id string) (GroupInfo, bool) {
	resp, ok := query.Request("groups.info",
		config.MakeURLValues(map[string]string{"channel": id}))
	if !ok || !resp.Ok {
		return GroupInfo{}, false
	}
	c, ok := parseInfo(*resp.Group)
	if !ok {
		return GroupInfo{}, false
	}
	return c, true
}

func parseInfo(data []byte) (GroupInfo, bool) {
	var g GroupInfo
	err := json.Unmarshal(data, &g)
	if err != nil {
		glog.Errorf("groups.info: Error parsing json: %v", err)
		return GroupInfo{}, false
	}
	return g, true
}

func (g GroupInfo) String() string {
	return pretty.Sprint(g)
}
