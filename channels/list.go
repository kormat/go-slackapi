package channels

import (
	"bitbucket.org/kormaton/slapi/config"
	"bitbucket.org/kormaton/slapi/query"
	"encoding/json"
	"github.com/golang/glog"
)

type ChannelList struct {
	Channels []json.RawMessage
}

func parseList(data []byte) (ChannelList, bool) {
	var l ChannelList
	err := json.Unmarshal(data, &l.Channels)
	if err != nil {
		glog.Errorf("channels.list: Error parsing json: %v", err)
		return ChannelList{}, false
	}
	return l, true
}

func List() ([]ChannelInfo, bool) {
	resp, ok := query.Request("channels.list", config.MakeURLValues(map[string]string{}))
	if !ok || !resp.Ok {
		return []ChannelInfo{}, false
	}
	cl, ok := parseList(*resp.Channels)
	if !ok {
		return []ChannelInfo{}, false
	}
	var infos []ChannelInfo
	for i, rawInfo := range cl.Channels {
		c, ok := parseInfo(rawInfo)
		if !ok {
			glog.Errorf("Error parsing channel %d", i)
			return []ChannelInfo{}, false
		}
		infos = append(infos, c)
	}
	return infos, true
}
