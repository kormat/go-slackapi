package chat

import (
	"fmt"
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/query"
)

func PostMsg(user, msg string) error {
	user = fmt.Sprintf("@%s", user)
	_, err := query.Request("chat.postMessage",
		config.MakeURLValues(map[string]string{"channel": user, "text": msg}))
	return err
}
