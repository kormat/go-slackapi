package users

import (
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/query"
	"github.com/kormat/go-slackapi/util"
	"github.com/kr/pretty"
)

type UserInfo struct {
	ID                string
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

func Info(id string) (UserInfo, error) {
	resp, err := query.Request("users.info",
		config.MakeURLValues(map[string]string{"user": id}))
	if err != nil {
		return UserInfo{}, err
	}
	return parseInfo(*resp.User)
}

func parseInfo(raw []byte) (UserInfo, error) {
	var u UserInfo
	if err := util.ParseJSON(raw, &u); err != nil {
		return UserInfo{}, util.ErrorLog("users.info: %v", err)
	}
	return u, nil
}

func (u UserInfo) String() string {
	return pretty.Sprint(u)
}
