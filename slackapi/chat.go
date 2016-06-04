package main

import (
	"github.com/kormat/go-slackapi/chat"
	"github.com/kormat/go-slackapi/config"
)

type PostMsg struct {
	Args struct {
		User string `description:"Username"`
		Msg  string `description:"Message to send"`
	} `positional-args:"yes" required:"yes"`
}

var postMsg PostMsg

func init() {
	parser.AddCommand("chat.postmsg", "Post message to user", "", &postMsg)
}

func (p *PostMsg) Execute(_ []string) error {
	if config.CfgErr != nil {
		return config.CfgErr
	}
	return chat.PostMsg(p.Args.User, p.Args.Msg)
}
