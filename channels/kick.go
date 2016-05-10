package channels

import (
	"bitbucket.org/kormaton/slapi/config"
	"bitbucket.org/kormaton/slapi/query"
)

func Kick(channel, user string) bool {
	r, ok := query.Request("channels.kick",
		config.MakeURLValues(map[string]string{"channel": channel, "user": user}))
	if !ok || !r.Ok {
		return false
	}
	return true
}
