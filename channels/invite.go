package channels

import (
	"bitbucket.org/kormaton/slapi/config"
	"bitbucket.org/kormaton/slapi/query"
)

func Invite(channel, user string) bool {
	r, ok := query.Request("channels.invite",
		config.MakeURLValues(map[string]string{"channel": channel, "user": user}))
	if !ok || !r.Ok {
		return false
	}
	return true
}
