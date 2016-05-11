package main

import (
	"errors"
	"fmt"
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/users/admin"
)

type UserInvite struct {
	Args struct {
		Email string `positional-arg-name:"EMAIL" description:"User's email address"`
	} `positional-args:"yes" required:"yes"`
}

var userInvite UserInvite

func init() {
	parser.AddCommand("users.admin.invite", "Invite user to slack server", "", &userInvite)
}

func (ui *UserInvite) Execute(_ []string) error {
	if config.CfgErr != nil {
		return config.CfgErr
	}
	ok := uadmin.Invite(ui.Args.Email)
	if !ok {
		return errors.New("users.admin.invite failure")
	}
	fmt.Println("Success.")
	return nil
}
