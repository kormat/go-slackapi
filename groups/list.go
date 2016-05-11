package groups

import (
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/query"
	"github.com/kormat/go-slackapi/util"
)

func List() ([]GroupInfo, error) {
	resp, err := query.Request("groups.list", config.MakeURLValues(map[string]string{}))
	if err != nil {
		return []GroupInfo{}, err
	}
	groups, err := util.ParseJSONList(*resp.Groups)
	if err != nil {
		return []GroupInfo{}, util.Error("groups.list: %v", err)
	}
	var infos []GroupInfo
	for i, rawInfo := range groups {
		c, err := parseInfo(rawInfo)
		if err != nil {
			return []GroupInfo{}, util.Error("Error parsing group %d", i)
		}
		infos = append(infos, c)
	}
	return infos, nil
}
