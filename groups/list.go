package groups

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/query"
)

type GroupList struct {
	Groups []json.RawMessage
}

func parseList(data []byte) (GroupList, bool) {
	var l GroupList
	err := json.Unmarshal(data, &l.Groups)
	if err != nil {
		glog.Errorf("groups.list: Error parsing json: %v", err)
		return GroupList{}, false
	}
	return l, true
}

func List() ([]GroupInfo, bool) {
	resp, ok := query.Request("groups.list", config.MakeURLValues(map[string]string{}))
	if !ok || !resp.Ok {
		return []GroupInfo{}, false
	}
	ul, ok := parseList(*resp.Groups)
	if !ok {
		return []GroupInfo{}, false
	}
	var infos []GroupInfo
	for i, rawInfo := range ul.Groups {
		c, ok := parseInfo(rawInfo)
		if !ok {
			glog.Errorf("Error parsing group %d", i)
			return []GroupInfo{}, false
		}
		infos = append(infos, c)
	}
	return infos, true
}
