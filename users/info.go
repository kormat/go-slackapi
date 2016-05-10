package users

import (
	"bitbucket.org/kormaton/slapi/config"
	"bitbucket.org/kormaton/slapi/query"
	"encoding/json"
	"github.com/golang/glog"
	"github.com/kr/pretty"
)

type UserInfo struct {
	Id                string
	TeamID            string `json:"team_id"`
	Name              string
	Deleted           bool
	Status            string
	Color             string
	RealName          string `json:"real_name"`
	TZ                string
	TZLabel           string `json:"tz_label"`
	TZOffset          int    `json:"tz_offset"`
	Profile           userProfile
	IsAdmin           bool `json:"is_admin"`
	IsOwner           bool `json:"is_owner"`
	IsPrimaryOwner    bool `json:"is_primary_owner"`
	IsRestricted      bool `json:"is_restricted"`
	IsUltraRestricted bool `json:"is_ultra_restricted"`
	IsBot             bool `json:"is_bot"`
	Has2fa            bool `json:"has_2fa"`
}

type userProfile struct {
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	AvatarHash         string `json:"avatar_hash"`
	RealName           string `json:"real_name"`
	RealNameNormalized string `json:"real_name_normalized"`
	Email              string
}

func Info(id string) (UserInfo, bool) {
	resp, ok := query.Request("users.info",
		config.MakeURLValues(map[string]string{"user": id}))
	if !ok || !resp.Ok {
		return UserInfo{}, false
	}
	u, ok := parseInfo(*resp.User)
	if !ok {
		return UserInfo{}, false
	}
	return u, true
}

func parseInfo(data []byte) (UserInfo, bool) {
	var u UserInfo
	err := json.Unmarshal(data, &u)
	if err != nil {
		glog.Errorf("users.info: Error parsing json: %v", err)
		return UserInfo{}, false
	}
	return u, true
}

func (u UserInfo) String() string {
	return pretty.Sprint(u)
}
