package groups

import (
	"github.com/kormat/go-slackapi/channels"
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/query"
	"github.com/kormat/go-slackapi/util"
	"github.com/kr/pretty"
)

type GroupInfo struct {
	ID         string
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

func Info(id string) (GroupInfo, error) {
	resp, err := query.Request("groups.info",
		config.MakeURLValues(map[string]string{"channel": id}))
	if err != nil {
		return GroupInfo{}, err
	}
	return parseInfo(*resp.Group)
}

func parseInfo(raw []byte) (GroupInfo, error) {
	var g GroupInfo
	if err := util.ParseJSON(raw, &g); err != nil {
		return GroupInfo{}, util.Error("groups.info: %v", err)
	}
	return g, nil
}

func (g GroupInfo) String() string {
	return pretty.Sprint(g)
}
