package chat

import (
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/query"
)

func PostMsg(target, msg string) error {
	_, err := query.Request("chat.postMessage",
		config.MakeURLValues(map[string]string{"channel": target, "text": msg}))
	return err
}
