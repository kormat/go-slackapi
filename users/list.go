package users

import (
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/query"
	"github.com/kormat/go-slackapi/util"
)

func List() ([]UserInfo, error) {
	resp, err := query.Request("users.list", config.MakeURLValues(map[string]string{}))
	if err != nil {
		return []UserInfo{}, err
	}
	users, err := util.ParseJSONList(*resp.Members)
	if err != nil {
		return []UserInfo{}, util.Error("users.list: %v", err)
	}
	var infos []UserInfo
	for i, rawInfo := range users {
		u, err := parseInfo(rawInfo)
		if err != nil {
			return []UserInfo{}, util.Error("Error parsing user %d", i)
		}
		infos = append(infos, u)
	}
	return infos, nil
}
