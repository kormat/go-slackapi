package channels

import (
	"bitbucket.org/kormaton/slapi/config"
	"bitbucket.org/kormaton/slapi/query"
	"encoding/json"
	"github.com/golang/glog"
	"github.com/kr/pretty"
	"net/url"
)

type ChannelInfo struct {
	Id                 string
	Name               string
	IsChannel          bool `json:"is_channel"`
	Created            int
	Creator            string
	IsArchived         bool `json:"is_archived"`
	IsGeneral          bool `json:"is_general"`
	Members            []string
	Topic              topicPurpose
	Purpose            topicPurpose
	IsMember           bool   `json:"is_member"`
	LastRead           string `json:"last_read"`
	Latest             latest
	UnreadCount        int `json:"unread_count"`
	UnreadCountDisplay int `json:"unread_count_display"`
}

type topicPurpose struct {
	Value   string
	Creator string
	LastSet int `json:"last_set"`
}

type latest struct {
	Type string
	User string
	Text string
	TS   string
}

func parseInfo(data []byte) (ChannelInfo, bool) {
	var c ChannelInfo
	err := json.Unmarshal(data, &c)
	if err != nil {
		glog.Errorf("channels.info: Error parsing json: %v", err)
		return ChannelInfo{}, false
	}
	return c, true
}

func (c ChannelInfo) String() string {
	return pretty.Sprint(c)
}

func Info(id string) (ChannelInfo, bool) {
	cfg := config.Load()
	v := url.Values{}
	v.Set("token", cfg.Token)
	v.Set("channel", id)
	resp, ok := query.Request("channels.info", v)
	if !ok || !resp.Ok {
		return ChannelInfo{}, false
	}
	c, ok := parseInfo(*resp.Channel)
	if !ok {
		return ChannelInfo{}, false
	}
	return c, true
}
