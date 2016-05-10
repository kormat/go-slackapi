package groups

import (
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/query"
)

func Invite(channel, user string) bool {
	r, ok := query.Request("groups.invite",
		config.MakeURLValues(map[string]string{"channel": channel, "user": user}))
	if !ok || !r.Ok {
		return false
	}
	return true
}
