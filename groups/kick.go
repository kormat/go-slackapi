package groups

import (
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/query"
)

func Kick(channel, user string) error {
	_, err := query.Request("groups.kick",
		config.MakeURLValues(map[string]string{"channel": channel, "user": user}))
	return err
}
