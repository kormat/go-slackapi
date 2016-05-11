package channels

import (
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/query"
	"github.com/kormat/go-slackapi/util"
)

func List() ([]ChannelInfo, error) {
	resp, err := query.Request("channels.list", config.MakeURLValues(map[string]string{}))
	if err != nil {
		return []ChannelInfo{}, err
	}
	chans, err := util.ParseJSONList(*resp.Channels)
	if err != nil {
		return []ChannelInfo{}, util.Error("channels.list: %v", err)
	}
	var infos []ChannelInfo
	for i, rawInfo := range chans {
		c, err := parseInfo(rawInfo)
		if err != nil {
			return []ChannelInfo{}, util.Error("Error parsing channel %d", i)
		}
		infos = append(infos, c)
	}
	return infos, nil
}
