package uadmin

/*
N.B. this api is undocumented "because it's still under development":
https://github.com/slackhq/slack-api-docs/issues/30#issuecomment-137582387
*/

import (
	"bitbucket.org/kormaton/slapi/config"
	"bitbucket.org/kormaton/slapi/query"
)

func Invite(email string) bool {
	r, ok := query.Request("users.admin.invite",
		config.MakeURLValues(map[string]string{"email": email}))
	if !ok || !r.Ok {
		return false
	}
	return true
}
