package users

import (
	"bitbucket.org/kormaton/slapi/config"
	"bitbucket.org/kormaton/slapi/query"
	"encoding/json"
	"github.com/golang/glog"
)

type UserList struct {
	Users []json.RawMessage
}

func parseList(data []byte) (UserList, bool) {
	var l UserList
	err := json.Unmarshal(data, &l.Users)
	if err != nil {
		glog.Errorf("users.list: Error parsing json: %v", err)
		return UserList{}, false
	}
	return l, true
}

func List() ([]UserInfo, bool) {
	resp, ok := query.Request("users.list", config.MakeURLValues(map[string]string{}))
	if !ok || !resp.Ok {
		return []UserInfo{}, false
	}
	ul, ok := parseList(*resp.Members)
	if !ok {
		return []UserInfo{}, false
	}
	var infos []UserInfo
	for i, rawInfo := range ul.Users {
		c, ok := parseInfo(rawInfo)
		if !ok {
			glog.Errorf("Error parsing channel %d", i)
			return []UserInfo{}, false
		}
		infos = append(infos, c)
	}
	return infos, true
}
