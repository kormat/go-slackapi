package groups

import (
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/query"
)

func Invite(channel, user string) error {
	_, err := query.Request("groups.invite",
		config.MakeURLValues(map[string]string{"channel": channel, "user": user}))
	return err
}
