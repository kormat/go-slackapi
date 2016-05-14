package channels

import (
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/query"
	"github.com/kormat/go-slackapi/util"
	"github.com/kr/pretty"
)

type ChannelInfo struct {
	ID                 string
	Name               string
	IsChannel          bool `json:"is_channel"`
	Created            int
	Creator            string
	IsArchived         bool `json:"is_archived"`
	IsGeneral          bool `json:"is_general"`
	Members            []string
	Topic              TopicPurpose
	Purpose            TopicPurpose
	IsMember           bool    `json:"is_member"`
	LastRead           float64 `json:"last_read,string"`
	Latest             latest
	UnreadCount        int `json:"unread_count"`
	UnreadCountDisplay int `json:"unread_count_display"`
}

type TopicPurpose struct {
	Value   string
	Creator string
	LastSet int `json:"last_set"`
}

type latest struct {
	Type string
	User string
	Text string
	TS   float64 `json:",string"`
}

func Info(id string) (ChannelInfo, error) {
	resp, err := query.Request("channels.info",
		config.MakeURLValues(map[string]string{"channel": id}))
	if err != nil {
		return ChannelInfo{}, err
	}
	return parseInfo(*resp.Channel)
}

func parseInfo(raw []byte) (ChannelInfo, error) {
	var c ChannelInfo
	if err := util.ParseJSON(raw, &c); err != nil {
		return ChannelInfo{}, util.Error("channels.info: %v", err)
	}
	return c, nil
}

func (c ChannelInfo) String() string {
	return pretty.Sprint(c)
}
